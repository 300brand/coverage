package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"time"
)

type Clock struct {
	Log     skynet.SemanticLogger
	Tickers map[string]*Ticker
}

type Ticker struct {
	F      func()
	Once   chan bool
	Start  chan bool
	Stop   chan bool
	Tick   time.Duration
	Ticker *time.Ticker
}

var _ service.ServiceDelegate = &Clock{}

func (s *Clock) Registered(service *service.Service) {
	s.Log.Trace("Registered")
	for name, t := range s.Tickers {
		s.Log.Trace("Starting " + name)
		go runner(t)
		t.Start <- true
	}
}
func (s *Clock) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}
func (s *Clock) Started(service *service.Service) {
	s.Log.Trace("Started")
	s.Tickers = map[string]*Ticker{}
	s.Tickers["QueueFeedAdder"] = &Ticker{
		F:      s.queueFeedAdder,
		Once:   make(chan bool, 1),
		Start:  make(chan bool, 1),
		Stop:   make(chan bool, 1),
		Tick:   time.Second * 10,
		Ticker: &time.Ticker{},
	}
}
func (s *Clock) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
	for name, t := range s.Tickers {
		s.Log.Trace("Stopping " + name)
		t.Stop <- true
	}
}
func (s *Clock) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}
func (s *Clock) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
