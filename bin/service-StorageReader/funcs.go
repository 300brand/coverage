package main

import (
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
)

func (s *StorageReader) OldestFeed(ri *skynet.RequestInfo, in *skytypes.ObjectIds, out *coverage.Feed) (err error) {
	s.Log.Debug(fmt.Sprintf("OldestFeed.in: %+v", in))
	oldest, err := s.Mongo.GetOldestFeed(in.Ids)
	*out = *oldest
	s.Log.Debug(fmt.Sprintf("OldestFeed.out: %+v", out))
	return
}
