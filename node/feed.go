package node

import (
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"labix.org/v2/mgo/bson"
)

type SkynetFeed struct {
	Log skynet.SemanticLogger
}

var _ service.ServiceDelegate = &SkynetFeed{}

func (s *SkynetFeed) Registered(service *service.Service)   { s.Log.Trace("Registered") }
func (s *SkynetFeed) Unregistered(service *service.Service) { s.Log.Trace("Unregistered") }
func (s *SkynetFeed) Started(service *service.Service)      { s.Log.Trace("Started") }
func (s *SkynetFeed) Stopped(service *service.Service)      { s.Log.Trace("Stopped") }
func (s *SkynetFeed) MethodCalled(method string)            { s.Log.Trace("MethodCalled") }
func (s *SkynetFeed) MethodCompleted(method string, duration int64, err error) {
	s.Log.Trace("MethodCompleted")
}

// in will be nil
func (s *SkynetFeed) Download(r *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	s.Log.Debug(fmt.Sprintf("%+v", in))
	s.Log.Debug(fmt.Sprintf("%+v", out))
	*out = *in
	return downloader.NewFeedService().Update(out)
}

func (s *SkynetFeed) NextFeed(r *skynet.RequestInfo, in map[bson.ObjectId]bool, out *coverage.Feed) (err error) {
	m := mongo.New("localhost", "Coverage")
	if err = m.Connect(); err != nil {
		return
	}
	defer m.Close()
	ignore := []bson.ObjectId{}
	f, err := m.GetOldestFeed(ignore)
	*out = *f
	return
}
