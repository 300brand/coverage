package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
	"time"
)

type Manager struct {
	Log             skynet.SemanticLogger
	Tickers         map[string]*Ticker
	ArticleBody     *client.ServiceClient
	ArticleDownload *client.ServiceClient
	ArticleLexer    *client.ServiceClient
	Feed            *client.ServiceClient
	FeedDownload    *client.ServiceClient
	FeedProcess     *client.ServiceClient
	Queue           *client.ServiceClient
	StorageWriter   *client.ServiceClient
}

var _ service.ServiceDelegate = &Manager{}

func (s *Manager) Registered(service *service.Service) {
	s.Log.Trace("Registered")

	s.ArticleBody = c.GetService("ArticleBody", "", "", "")
	s.ArticleDownload = c.GetService("ArticleDownload", "", "", "")
	s.ArticleDownload.SetTimeout(0, 60*time.Second)
	s.ArticleLexer = c.GetService("ArticleLexer", "", "", "")
	s.Feed = c.GetService("Feed", "", "", "")
	s.FeedDownload = c.GetService("FeedDownload", "", "", "")
	s.FeedDownload.SetTimeout(0, 60*time.Second)
	s.FeedProcess = c.GetService("FeedProcess", "", "", "")
	s.FeedProcess.SetTimeout(0, 30*time.Second)
	s.Queue = c.GetService("Queue", "", "", "")
	s.StorageWriter = c.GetService("StorageWriter", "", "", "")

	for name, t := range s.Tickers {
		s.Log.Trace("Starting " + name)
		go runner(t)
		// if autoStart {
		// 	t.Start <- true
		// }
	}
}
func (s *Manager) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}
func (s *Manager) Started(service *service.Service) {
	s.Log.Trace("Started")
	s.Tickers = map[string]*Ticker{
		"QueueFeedAdder": NewTicker(s.queueFeedAdder, time.Second*10),
		"FeedProcessor":  NewTicker(s.feedProcessor, time.Second*10),
	}
}
func (s *Manager) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
	for name, t := range s.Tickers {
		s.Log.Trace("Stopping " + name)
		t.Stop <- true
	}
}
func (s *Manager) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}
func (s *Manager) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
