package node

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type FeedDownload struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &FeedDownload{}

func (s *FeedDownload) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *FeedDownload) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *FeedDownload) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *FeedDownload) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *FeedDownload) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *FeedDownload) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
