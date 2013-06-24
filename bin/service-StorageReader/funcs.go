package main

import (
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
)

func (s *StorageReader) GetFeed(ri *skynet.RequestInfo, in *skytypes.ObjectId, out *coverage.Feed) (err error) {
	s.Log.Debug(fmt.Sprintf("GetFeed.in: %+v", in.Id))
	f, err := s.Mongo.GetFeed(in.Id)
	*out = *f
	s.Log.Debug(fmt.Sprintf("GetFeed.out: %+v", out.ID))
	return
}

func (s *StorageReader) OldestFeed(ri *skynet.RequestInfo, in *skytypes.ObjectIds, out *coverage.Feed) (err error) {
	s.Log.Debug(fmt.Sprintf("OldestFeed.in: %+v", in))
	oldest, err := s.Mongo.GetOldestFeed(in.Ids)
	*out = *oldest
	s.Log.Debug(fmt.Sprintf("OldestFeed.out: %+v", out.ID))
	return
}
