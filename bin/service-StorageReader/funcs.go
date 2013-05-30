package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
	"labix.org/v2/mgo/bson"
)

func (s *StorageReader) OldestFeed(ri *skynet.RequestInfo, in []bson.ObjectId, out *coverage.Feed) (err error) {
	out, err = s.Mongo.GetOldestFeed(in)
	return
}
