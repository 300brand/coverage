package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/skynet"
	"labix.org/v2/mgo/bson"
)

const FeedQueue = "FeedQueue"

func (s *StorageReader) QueueFeedIds(ri *skynet.RequestInfo, in interface{}, out []bson.ObjectId) (err error) {
	c, err := s.Mongo.CappedCollection("FeedQueue", mongo.DefaultCapacity)
	if err != nil {
		return
	}
	out, err = c.AllIds()
	return
}

func (s *StorageReader) OldestFeed(ri *skynet.RequestInfo, in []bson.ObjectId, out *coverage.Feed) (err error) {
	out, err = s.Mongo.GetOldestFeed(in)
	return
}
