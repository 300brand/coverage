package node

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type Queue struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &Queue{}

func (s *Queue) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *Queue) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *Queue) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *Queue) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *Queue) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *Queue) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
