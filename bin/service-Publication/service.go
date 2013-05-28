package main

import (
	"errors"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"strings"
)

type Publication struct {
	Log       skynet.SemanticLogger
	MongoHost string
	MongoDb   string
	m         *mongo.Mongo
}

var _ service.ServiceDelegate = &Publication{}

func (s *Publication) Registered(service *service.Service) {
	s.Log.Trace("Registered")
}

func (s *Publication) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}

func (s *Publication) Started(service *service.Service) {
	s.Log.Trace("Started")
	s.m = mongo.New(s.MongoHost, s.MongoDb)
	s.Log.Trace("Connecting to MongoDB: " + s.MongoHost + " " + s.MongoDb)
	if err := s.m.Connect(); err != nil {
		s.Log.Fatal(err.Error())
	}
	s.Log.Trace("Connected to MongoDB")
	if err := s.m.EnsureIndexes(); err != nil {
		s.Log.Fatal(err.Error())
	}
	s.Log.Trace("Ensured indexes")
}

func (s *Publication) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
	s.m.Close()
	s.Log.Trace("Closed connection to MongoDB")
}

func (s *Publication) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}

func (s *Publication) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}


