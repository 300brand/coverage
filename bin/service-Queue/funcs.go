package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
	"labix.org/v2/mgo/bson"
)

func (s *Queue) AddFeed(ri *skynet.RequestInfo, in *bson.ObjectId, out interface{}) (err error) {
	feed := &coverage.Feed{}
	if err = c.GetService("StorageReader", "", "", "").Send(ri, "OldestFeed", in, feed); err != nil {
		return
	}
	if err = s.FeedQ.Push(feed.ID); err != nil {
		return
	}
	return
}
