package main

import (
	"git.300brand.com/coverage/doozer/idqueue"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
)

type Queue struct {
	Log   skynet.SemanticLogger
	FeedQ *idqueue.IdQueue
	Feed  *client.ServiceClient
}

var _ service.ServiceDelegate = &Queue{}

func (s *Queue) Registered(service *service.Service) {
	s.Log.Trace("Registered")
	s.Feed = c.GetService("Feed", "", "", "")
}
func (s *Queue) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}
func (s *Queue) Started(service *service.Service) {
	s.Log.Trace("Started")
	if err := s.FeedQ.Connect(); err != nil {
		s.Log.Fatal(err.Error())
	}
	s.Log.Trace("Connected to doozer [" + s.FeedQ.Addr + "] for queue storage")
}
func (s *Queue) Stopped(service *service.Service) {
	s.Log.Trace("Closing connection to doozer for queue storage")
	s.FeedQ.Close()
	s.Log.Trace("Stopped")
}
func (s *Queue) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}
func (s *Queue) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
