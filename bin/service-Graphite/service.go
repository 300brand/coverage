package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type Graphite struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &Graphite{}

func (s *Graphite) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *Graphite) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *Graphite) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *Graphite) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *Graphite) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *Graphite) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
