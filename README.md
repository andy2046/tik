# tik

[![Documentation](https://godoc.org/github.com/andy2046/tik?status.svg)](http://godoc.org/github.com/andy2046/tik)
[![GitHub issues](https://img.shields.io/github/issues/andy2046/tik.svg)](https://github.com/andy2046/tik/issues)
[![license](https://img.shields.io/github/license/andy2046/tik.svg)](https://github.com/andy2046/tik/LICENSE)
[![Release](https://img.shields.io/github/release/andy2046/tik.svg?label=Release)](https://github.com/andy2046/tik/releases)

----

## hierarchical timing wheel made easy

simplified version of [timeout](https://github.com/wahern/timeout) in Golang

for documentation, view the [API reference](./doc.md)


## Install

```
go get github.com/andy2046/tik
```

## Usage

```go
package main

import (
	"sync"
	"time"

	"github.com/andy2046/tik"
)

func main() {
	var l sync.RWMutex
	// init a new instance
	tk := tik.New()
	i := 0
	cb := func() {
		l.Lock()
		i++
		l.Unlock()
	}
	// schedule to run cb in 500ms
	to := tk.Schedule(500, cb)

	if !to.Pending() {
		panic("it should be pending")
	}

	if to.Expired() {
		panic("it should NOT be expired")
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
		panic("fail to callback", i)
	}
}
```
