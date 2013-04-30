package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type ArticleLexer struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &ArticleLexer{}

func (s *ArticleLexer) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *ArticleLexer) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *ArticleLexer) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *ArticleLexer) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *ArticleLexer) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *ArticleLexer) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
