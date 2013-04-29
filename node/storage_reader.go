package node

import (
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type StorageReader struct {
	Log   skynet.SemanticLogger
	Mongo *mongo.Mongo
}

var _ service.ServiceDelegate = &StorageReader{}

func (s *StorageReader) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *StorageReader) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *StorageReader) Started(service *service.Service) {
	s.Log.Trace("Started")
}
func (s *StorageReader) Stopped(service *service.Service) { s.Log.Trace("Stopped") }
func (s *StorageReader) MethodCalled(method string)       { s.Log.Trace("MethodCalled") }
func (s *StorageReader) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
