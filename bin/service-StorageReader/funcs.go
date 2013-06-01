package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
)

func (s *StorageReader) OldestFeed(ri *skynet.RequestInfo, in *skytypes.ObjectIds, out *coverage.Feed) (err error) {
	out, err = s.Mongo.GetOldestFeed(in.Ids)
	return
}
