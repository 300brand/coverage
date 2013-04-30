package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type ArticleBody struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &ArticleBody{}

func (s *ArticleBody) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *ArticleBody) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *ArticleBody) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *ArticleBody) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *ArticleBody) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *ArticleBody) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
