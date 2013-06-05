package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"time"
)

type Manager struct {
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

var _ service.ServiceDelegate = &Manager{}

func (s *Manager) Registered(service *service.Service) {
	s.Log.Trace("Registered")

	SCQueue = c.GetService("Queue", "", "", "")

	for name, t := range s.Tickers {
		s.Log.Trace("Starting " + name)
		go runner(t)
		t.Start <- true
	}
}
func (s *Manager) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}
func (s *Manager) Started(service *service.Service) {
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
func (s *Manager) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
	for name, t := range s.Tickers {
		s.Log.Trace("Stopping " + name)
		t.Stop <- true
	}
}
func (s *Manager) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}
func (s *Manager) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
