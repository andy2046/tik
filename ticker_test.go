package tik

import (
	"sync"
	"testing"
	"time"
)

func TestOneShotRun(t *testing.T) {
	tk := New()
	var l sync.RWMutex
	i := 0
	cb := func() {
		l.Lock()
		i++
		l.Unlock()
	}

	to := tk.Schedule(500, cb)

	if !to.Pending() {
		t.Error("it should be pending")
	}

	if to.Expired() {
		t.Error("it should NOT be expired")
	}

	for {
		time.Sleep(100 * time.Millisecond)

		if tk.AnyPending() {
			continue
		}

		if tk.AnyExpired() {
			continue
		}

		break
	}

	l.RLock()
	defer l.RUnlock()

	if i != 1 {
		t.Error("fail to callback", i)
	}
}

func TestOneShotCancel(t *testing.T) {
	tk := New()
	i := 0
	cb := func() {
		i++
	}

	to := tk.Schedule(500, cb)

	if !to.Pending() {
		t.Error("it should be pending")
	}

	if to.Expired() {
		t.Error("it should NOT be expired")
	}

	go func() {
		time.Sleep(200 * time.Millisecond)
		tk.Cancel(to)
	}()

	for {
		time.Sleep(100 * time.Millisecond)

		if tk.AnyPending() {
			continue
		}

		if tk.AnyExpired() {
			continue
		}

		break
	}

	if i != 0 {
		t.Error("fail to Cancel", i)
	}
}

func TestCancelAndCheck(t *testing.T) {
	tk := New()
	i := 0
	cb := func() {
		i++
	}

	to := tk.Schedule(500, cb)

	go func() {
		for i := 0; i < 10; i++ {
			to.Pending()
			to.Expired()
		}
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		tk.Cancel(to)
	}()

	for {
		time.Sleep(100 * time.Millisecond)

		if tk.AnyPending() {
			continue
		}

		if tk.AnyExpired() {
			continue
		}

		break
	}

	if i != 0 {
		t.Error("fail to Cancel", i)
	}
}

func TestRunAndCancel(t *testing.T) {
	var l sync.RWMutex
	var to *Timeout

	tk := New()
	i := 0
	n := 10
	cb := func() {
		l.Lock()
		i++
		l.Unlock()
	}

	go func() {
		for j := 0; j < n; j++ {
			to = tk.Schedule(500, cb)

			go func(to *Timeout) {
				for i := 0; i < 10; i++ {
					to.Pending()
					to.Expired()
				}
			}(to)

			if 0x1&j == 1 {
				go func(to *Timeout) {
					tk.Cancel(to)
				}(to)
			}
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)

		if tk.AnyPending() {
			continue
		}

		if tk.AnyExpired() {
			continue
		}

		break
	}

	l.RLock()
	defer l.RUnlock()

	if i != 5 {
		t.Error("fail to run", i)
	}
}

func TestRunPanic(t *testing.T) {
	tk := New()
	cb := func() {
		panic("test")
	}

	to := tk.Schedule(100, cb)

	if !to.Pending() {
		t.Error("it should be pending")
	}

	if to.Expired() {
		t.Error("it should NOT be expired")
	}

	for {
		time.Sleep(100 * time.Millisecond)

		if tk.AnyPending() {
			continue
		}

		if tk.AnyExpired() {
			continue
		}

		break
	}
}

func TestClose(t *testing.T) {
	tk := New()
	cb := func() {}

	tk.Close()
	to := tk.Schedule(500, cb)
	tk.Cancel(to)

	if to != nil {
		t.Error("it should be nil")
	}
}
