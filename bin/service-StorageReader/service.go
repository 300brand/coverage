package main

import (
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type StorageReader struct {
	Log       skynet.SemanticLogger
	Mongo     *mongo.Mongo
	MongoHost string
	MongoDb   string
}

var _ service.ServiceDelegate = &StorageReader{}

func (s *StorageReader) Registered(service *service.Service) {
	s.Log.Trace("Registered")
}

func (s *StorageReader) Unregistered(service *service.Service) {
	s.Log.Trace("Unregistered")
}

func (s *StorageReader) Started(service *service.Service) {
	s.Log.Trace("Started")
	s.Mongo = mongo.New(s.MongoHost, s.MongoDb)
	s.Log.Trace("Connecting to MongoDB: " + s.MongoHost + " " + s.MongoDb)
	if err := s.Mongo.Connect(); err != nil {
		s.Log.Fatal(err.Error())
	}
	s.Log.Trace("Connected to MongoDB")
	if err := s.Mongo.EnsureIndexes(); err != nil {
		s.Log.Fatal(err.Error())
	}
	s.Log.Trace("Ensured indexes")
}

func (s *StorageReader) Stopped(service *service.Service) {
	s.Log.Trace("Stopped")
	s.Mongo.Close()
	s.Log.Trace("Closed connection to MongoDB")
}

func (s *StorageReader) MethodCalled(method string) {
	s.Log.Trace("MethodCalled")
}

func (s *StorageReader) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
