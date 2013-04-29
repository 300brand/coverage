package node

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
)

type StorageWriter struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &StorageWriter{}

func (s *StorageWriter) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *StorageWriter) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *StorageWriter) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *StorageWriter) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *StorageWriter) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *StorageWriter) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}
