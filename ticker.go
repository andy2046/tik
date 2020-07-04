// Package tik implements Hierarchical Timing Wheels.
package tik

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	list "github.com/andy2046/gopie/pkg/dll"
	"github.com/andy2046/gopie/pkg/log"
	"github.com/andy2046/maths"
	"golang.org/x/sys/cpu"
)

const (
	cacheLinePadSize = unsafe.Sizeof(cpu.CacheLinePad{})
)

var (
	logger = log.NewLogger(func(c *log.Config) error {
		c.Level = log.DEBUG
		c.Prefix = "<tik>\t"
		return nil
	})
)

type (
	// Ticker progress the timing wheels.
	Ticker struct {
		_      [cacheLinePadSize]byte
		closed uint64
		_      [cacheLinePadSize - 8%cacheLinePadSize]byte

		curtime    uint64
		wheelLen   uint64 // 1<<nWheelBit
		wheelMax   uint64 // wheelLen-1
		wheelMask  uint64 // wheelLen-1
		timeoutMax uint64 // 1<<(nWheel*nWheelBit)-1

		// nWheel is the number of wheels. nWheelBit * nWheel is the number of
		// value bits used by all the wheels.
		nWheel uint8
		// nWheelBit is the number of value bits mapped in each wheel.
		// nWheelBit can NOT be larger than 6 bits because 2^6 -> 64 is the largest
		// number of slots which can be tracked.
		nWheelBit uint8

		closer  chan struct{}
		timer   Timer
		expired *list.List
		wheelz  []*wheel
		once    sync.Once
	}

	wheel struct {
		_       [cacheLinePadSize]byte
		pending uint64
		_       [cacheLinePadSize - 8%cacheLinePadSize]byte
		slot    []*list.List
	}

	// Config used to init Ticker.
	Config struct {
		// number of value bits mapped in each wheel.
		WheelBitNum uint8
		// number of wheels.
		WheelNum uint8
		// Timer to progress the timing wheel.
		Timer Timer
	}

	// Option applies config to Ticker Config.
	Option = func(*Config)
)

var (
	// DefaultConfig is the default Ticker Config.
	DefaultConfig = Config{
		WheelBitNum: 6,
		WheelNum:    4,
		Timer:       nil,
	}
)

// New initiates a new Ticker.
func New(options ...Option) *Ticker {
	tConfig := DefaultConfig
	setOption(&tConfig, options...)

	if tConfig.WheelBitNum < 3 || tConfig.WheelBitNum > 6 {
		panic("WheelBitNum should be in range [3, 6]")
	}

	switch tConfig.WheelNum {
	case 2:
	case 4:
	case 8:
	default:
		panic("WheelNum should be in set {2, 4, 8}")
	}

	if tConfig.Timer == nil {
		tConfig.Timer = NewTimer(100) // 100 millisecond interval
	}

	tk := &Ticker{}
	tk.init(tConfig.WheelBitNum, tConfig.WheelNum, tConfig.Timer)

	return tk
}

func (tk *Ticker) init(nWheelBit, nWheel uint8, timer Timer) {
	tk.nWheel = nWheel
	tk.nWheelBit = nWheelBit
	tk.wheelLen = 1 << nWheelBit
	tk.wheelMax = 1<<nWheelBit - 1
	tk.wheelMask = 1<<nWheelBit - 1
	tk.timeoutMax = 1<<(nWheel*nWheelBit) - 1

	tk.closer = make(chan struct{})
	tk.timer = timer
	tk.curtime = timer.Now()
	tk.expired = list.New()
	tk.wheelz = make([]*wheel, tk.nWheel)

	for i := range tk.wheelz {
		slot := make([]*list.List, tk.wheelLen)
		for j := range slot {
			slot[j] = list.New()
		}

		tk.wheelz[i] = &wheel{
			pending: 0,
			slot:    slot,
		}
	}

	go func() {
		for {
			select {
			case <-tk.closer:
				return
			case absolute := <-tk.timer.Step():
				tk.update(absolute)
			}
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				logger.Error(r)
				logger.Error(string(buf[:n]))
			}
		}()

		for {
			select {
			case <-tk.closer:
				return
			default:
			}

			to := tk.getExpired()
			if to != nil && to.callbk != nil {
				to.callbk()
			}
		}
	}()
}

// Schedule creates a one-shot action that executed after the given delay.
// `delay` is the time from now to delay execution,
// the time unit of the delay depends on the Timer provided, default is millisecond.
// `cb` is the task to execute.
// it returns `nil` if Ticker is closed.
func (tk *Ticker) Schedule(delay uint64, cb Callback) *Timeout {
	if tk.IsClosed() {
		logger.Error("Ticker closed, no more Schedule is allowed")
		return nil
	}

	to := newTimeout(cb)
	tk.add(to, delay)
	return to
}

// Cancel the Timeout scheduled if it has not yet expired.
func (tk *Ticker) Cancel(to *Timeout) {
	if to == nil {
		return
	}

	tk.del(to)
}

// Close stop processing any task,
// whether it is pending or expired.
func (tk *Ticker) Close() {
	tk.once.Do(func() {
		close(tk.closer)
	})

	atomic.CompareAndSwapUint64(&tk.closed, 0, 1)
	tk.timer.Stop()
}

// IsClosed returns true if closed, false otherwise.
func (tk *Ticker) IsClosed() bool {
	return atomic.LoadUint64(&tk.closed) == 1
}

func (tk *Ticker) update(curtime uint64) {
	curtimeTK := atomic.LoadUint64(&tk.curtime)
	elapsed := curtime - curtimeTK
	todo := list.New()

	// looping over every wheel, better to keep number of wheelz smallish
	for i := uint8(0); i < tk.nWheel; i++ {
		pending := uint64(0)
		iWheelBit := i * tk.nWheelBit

		// calculate the slots expiring in this wheel
		//
		//  if the elapsed time is greater than the maximum period of
		//  the wheel, mark every position as expiring.
		//
		//  otherwise, to determine the expired slots fill in all the
		//  bits between the last slot processed and the current
		//  slot, inclusive of the last slot.
		//
		//  if a wheel rolls over, force a tick of the next higher
		//  wheel.
		if (elapsed >> iWheelBit) > tk.wheelMax {
			pending = math.MaxUint64
		} else {
			elapsedTmp := tk.wheelMask & (elapsed >> iWheelBit)
			elapsedTmpShift := uint64((1 << elapsedTmp) - 1)

			oslot := tk.wheelMask & (curtimeTK >> iWheelBit)
			pending = rotl(elapsedTmpShift, int(oslot))

			nslot := tk.wheelMask & (curtime >> iWheelBit)
			pending |= rotr(rotl(elapsedTmpShift, int(nslot)), int(elapsedTmp))
			pending |= (1 << nslot)
		}

		pp := &tk.wheelz[i].pending
		for pending&atomic.LoadUint64(pp) != 0 {
			islot := ctz(pending & tk.wheelz[i].pending)
			l := tk.wheelz[i].slot[islot]
			for !l.Empty() {
				ee := l.PopLeft()
				if e, ok := ee.(*Timeout); ok {
					todo.PushRight(e)
				}
			}

			for l.Empty() {
				p := atomic.LoadUint64(pp)

				if casPending(pp, p, p&^(1<<uint(islot))) { // race
					break
				}
			}
		}

		if 0x1&pending == 0 {
			// break if not wrapping around end of wheel.
			break
		}

		// if continuing, the next wheel must tick at least once.
		elapsed = maths.Uint64Var.Max(elapsed, tk.wheelLen<<iWheelBit)
	}

	atomic.StoreUint64(&tk.curtime, curtime)

	for !todo.Empty() {
		to := todo.PopLeft().(*Timeout)
		to.m.Lock()
		to.pending = nil
		to.element = nil // unlink to List
		to.m.Unlock()
		tk.sched(to, to.expires, curtime)
	}
}

func (tk *Ticker) add(to *Timeout, timeout uint64) {
	curtime := atomic.LoadUint64(&tk.curtime)
	tk.sched(to, curtime+timeout, curtime)
}

func (tk *Ticker) sched(to *Timeout, expires, curtime uint64) {
	tk.del(to)

	if expires > curtime {
		rem := tk.rem(expires, curtime)
		iwheel := tk.wheel(rem)
		islot := tk.slot(iwheel, expires)

		to.m.Lock()
		to.iwheel, to.islot = iwheel, islot
		to.pending = tk.wheelz[iwheel].slot[islot]
		to.element = to.pending.PushRight(to) // link to List
		to.expires = expires
		to.papa = tk
		pending := to.pending
		to.m.Unlock()

		for !pending.Empty() {
			pp := &tk.wheelz[iwheel].pending
			p := atomic.LoadUint64(pp)

			if p&(1<<islot) != 0 {
				break
			}

			if casPending(pp, p, p|(1<<islot)) { // race
				break
			}
		}
	} else {
		to.m.Lock()
		to.expires = expires
		to.papa = tk
		to.pending = tk.expired
		to.element = to.pending.PushRight(to)
		to.m.Unlock()
	}
}

func (tk *Ticker) del(to *Timeout) {
	to.m.Lock()
	defer to.m.Unlock()

	if to.pending != nil {
		if to.element != nil {
			to.pending.Remove(to.element)
		}

		for {
			if to.pending != tk.expired && to.pending.Empty() &&
				to.iwheel != math.MaxUint8 && to.islot != math.MaxUint8 {
				pp := &tk.wheelz[to.iwheel].pending
				p := atomic.LoadUint64(pp)
				if casPending(pp, p, p&^(1<<to.islot)) { // race
					break
				}
			} else {
				break
			}
		}

		to.pending = nil
		to.papa = nil
		to.element = nil
	}
}

func (tk *Ticker) wheel(timeout uint64) uint8 {
	// must be timeout != 0, so fls input is nonzero
	return uint8(fls(maths.Uint64Var.Min(timeout, tk.timeoutMax))-1) / tk.nWheelBit
}

func (tk *Ticker) slot(wheel uint8, expires uint64) uint8 {
	i := uint64(1)
	if wheel == 0 {
		i = 0
	}
	return uint8(tk.wheelMask & ((expires >> (wheel * tk.nWheelBit)) - i))
}

func (tk *Ticker) rem(expires, curtime uint64) uint64 {
	return expires - curtime
}

func (tk *Ticker) getExpired() *Timeout {
	if tk.expired.Empty() {
		return nil
	}

	to := tk.expired.PopLeft().(*Timeout)
	to.m.Lock()
	to.pending = nil
	to.papa = nil
	to.element = nil
	to.m.Unlock()
	return to
}

/*
// calculate the interval before needing to process any timeouts pending on
// any wheel.
//
// this might return a timeout value sooner than any installed timeout if
// only higher-order wheels have timeouts pending. It only known when to
// process a wheel, not precisely when a timeout is scheduled.
func (tk *Ticker) interval() uint64 {
	var timeoutTmp, relmask, pending uint64
	var timeout uint64 = math.MaxUint64
	var curtime = atomic.LoadUint64(&tk.curtime)

	for i := uint8(0); i < tk.nWheel; i++ {
		pending = atomic.LoadUint64(&tk.wheelz[i].pending)
		if pending != 0 {
			iWheelBit := i * tk.nWheelBit
			slot := tk.wheelMask & (curtime >> iWheelBit)
			j := 1
			if i == 0 {
				j = 0
			}
			timeoutTmp = uint64(ctz(rotr(pending, int(slot)))+j) << iWheelBit
			timeoutTmp -= (relmask & curtime)

			timeout = maths.Uint64Var.Min(timeoutTmp, timeout)
		}

		relmask <<= tk.nWheelBit
		relmask |= tk.wheelMask
	}

	return timeout
}
*/

// AnyExpired returns true if expiry queue is not empty, false otherwise.
func (tk *Ticker) AnyExpired() bool {
	return !tk.expired.Empty()
}

// AnyPending returns true if there is task in wheels, false otherwise.
func (tk *Ticker) AnyPending() bool {
	var p uint64

	for _, w := range tk.wheelz {
		p |= atomic.LoadUint64(&w.pending) // race
	}

	return p != 0
}

// setOption takes one or more Option function and applies them in order to Ticker Config.
func setOption(p *Config, options ...func(*Config)) {
	for _, opt := range options {
		opt(p)
	}
}

func casPending(p *uint64, c, n uint64) bool {
	return atomic.CompareAndSwapUint64(p, c, n)
}
