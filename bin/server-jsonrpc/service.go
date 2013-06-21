package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type ServerJSONRPC struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &ServerJSONRPC{}

func (s *ServerJSONRPC) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *ServerJSONRPC) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *ServerJSONRPC) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *ServerJSONRPC) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *ServerJSONRPC) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *ServerJSONRPC) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
