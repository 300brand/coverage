package node

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type FeedArticle struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &FeedArticle{}

func (s *FeedArticle) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *FeedArticle) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *FeedArticle) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *FeedArticle) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *FeedArticle) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *FeedArticle) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
