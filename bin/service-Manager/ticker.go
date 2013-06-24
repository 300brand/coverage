package main

import (
	"sync/atomic"
	"time"
)

type Ticker struct {
	F       func() error
	Once    chan bool
	Running int64
	Start   chan bool
	Stop    chan bool
	Tick    time.Duration
	Ticker  *time.Ticker
}

func NewTicker(f func() error, d time.Duration) *Ticker {
	return &Ticker{
		F:      f,
		Once:   make(chan bool),
		Start:  make(chan bool),
		Stop:   make(chan bool),
		Tick:   d,
		Ticker: &time.Ticker{},
	}
}

func runner(t *Ticker) {
	for {
		select {
		case <-t.Once:
			// Moved to processCommand where error is returned
			// t.F()
		case <-t.Start:
			t.Ticker = time.NewTicker(t.Tick)
		case <-t.Stop:
			t.Ticker.Stop()
		case <-t.Ticker.C:
			go func(t *Ticker) {
				atomic.AddInt64(&t.Running, 1)
				t.F()
				atomic.AddInt64(&t.Running, -1)
			}(t)
		}
	}
}
