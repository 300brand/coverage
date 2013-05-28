package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
)

type Publication struct {
	Log    skynet.SemanticLogger
	Reader *client.ServiceClient
	Writer *client.ServiceClient
}

var _ service.ServiceDelegate = &Publication{}

func (s *Publication) Registered(service *service.Service) {
	s.Log.Trace("Registered")
}

func (s *Publication) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}

func (s *Publication) Started(service *service.Service) {
	s.Log.Trace("Started")
	s.Reader = c.GetService("StorageReader", "", "", "")
	s.Writer = c.GetService("StorageWriter", "", "", "")
}

func (s *Publication) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
}

func (s *Publication) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}

func (s *Publication) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
