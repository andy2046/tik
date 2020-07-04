package tik

import (
	"math"
	"sync"

	list "github.com/andy2046/gopie/pkg/dll"
)

type (
	// Callback function to trigger when timeout expires.
	Callback func()

	// Timeout represents user timeout logic.
	Timeout struct {
		// absolute expiration time
		expires uint64
		// callbk func when expires
		callbk Callback
		// timeout list if pending on wheel or expiry queue
		pending *list.List
		// pointer to ticker
		papa *Ticker

		element *list.Element

		iwheel uint8
		islot  uint8

		m sync.RWMutex
	}
)

func newTimeout(cb Callback) *Timeout {
	to := &Timeout{}
	to.init(cb)
	return to
}

// init initialize timeout.
func (to *Timeout) init(cb Callback) {
	to.callbk = cb
	to.iwheel = math.MaxUint8
	to.islot = math.MaxUint8
}

// Pending returns true if timeout is in timing wheel, false otherwise.
func (to *Timeout) Pending() bool {
	to.m.RLock()
	defer to.m.RUnlock()

	return to.pending != nil && to.papa != nil && to.pending != to.papa.expired
}

// Expired returns true if timeout is in expired queue, false otherwise.
func (to *Timeout) Expired() bool {
	to.m.RLock()
	defer to.m.RUnlock()

	return to.pending != nil && to.papa != nil && to.pending == to.papa.expired
}
