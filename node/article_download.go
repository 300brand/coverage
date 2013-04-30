package node

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type ArticleDownload struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &ArticleDownload{}

func (s *ArticleDownload) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *ArticleDownload) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *ArticleDownload) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *ArticleDownload) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *ArticleDownload) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *ArticleDownload) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
