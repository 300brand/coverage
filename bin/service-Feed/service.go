package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
)

type Feed struct {
	Log           skynet.SemanticLogger
	FeedDownload  *client.ServiceClient
	FeedProcess   *client.ServiceClient
	StorageReader *client.ServiceClient
	StorageWriter *client.ServiceClient
}

var _ service.ServiceDelegate = &Feed{}

func (s *Feed) Registered(service *service.Service) {
	s.Log.Trace("Registered")
	s.FeedDownload = c.GetService("FeedDownload", "", "", "")
	s.FeedProcess = c.GetService("FeedProcess", "", "", "")
	s.StorageReader = c.GetService("StorageReader", "", "", "")
	s.StorageWriter = c.GetService("StorageWriter", "", "", "")
}

func (s *Feed) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}

func (s *Feed) Started(service *service.Service) {
	s.Log.Trace("Started")
}

func (s *Feed) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
}

func (s *Feed) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}

func (s *Feed) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
