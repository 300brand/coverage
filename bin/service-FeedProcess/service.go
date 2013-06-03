package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type FeedProcess struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &FeedProcess{}

func (s *FeedProcess) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *FeedProcess) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *FeedProcess) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *FeedProcess) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *FeedProcess) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *FeedProcess) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
