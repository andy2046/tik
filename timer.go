package tik

import (
	"sync"
	"time"
)

type (
	// Timer progress the timing wheel by current time.
	Timer interface {
		// Now returns the absolute time when timer started.
		Now() uint64

		// Step channel to step timing wheel by absolute time.
		Step() <-chan uint64

		// Stop the timer.
		Stop()
	}

	// DefaultTimer implements Timer interface.
	DefaultTimer struct {
		nowTime  time.Time
		interval uint64
		absolute chan uint64
		closer   chan struct{}
		once     sync.Once
	}
)

// NewTimer returns DefaultTimer with interval in millisecond.
func NewTimer(interval uint64) Timer {
	dt := &DefaultTimer{
		nowTime:  time.Now(),
		interval: interval,
		absolute: make(chan uint64),
		closer:   make(chan struct{}),
	}

	go func() {
		tick := time.NewTicker(time.Duration(dt.interval) * time.Millisecond)
		defer tick.Stop()

		for {
			select {
			case <-dt.closer:
				return
			case t := <-tick.C:
				select {
				// Go>=1.9 uses a monotonic clock for duration
				case dt.absolute <- uint64(t.Sub(dt.nowTime) / time.Millisecond):
				default:
				}
			}
		}
	}()

	return dt
}

// Now returns the absolute time when timer started in millisecond.
func (dt *DefaultTimer) Now() uint64 {
	return 0
}

// Step timing wheel by absolute time in millisecond.
func (dt *DefaultTimer) Step() <-chan uint64 {
	return dt.absolute
}

// Stop the DefaultTimer.
func (dt *DefaultTimer) Stop() {
	dt.once.Do(func() {
		close(dt.closer)
	})
}
