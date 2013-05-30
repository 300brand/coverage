package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
	"labix.org/v2/mgo/bson"
)

func (s *Queue) AddFeed(ri *skynet.RequestInfo, in *interface{}, out *bson.ObjectId) (err error) {
	ids, err := s.FeedQ.Get()
	if err != nil {
		return
	}
	feed := &coverage.Feed{}
	if err = c.GetService("StorageReader", "", "", "").Send(ri, "OldestFeed", ids, feed); err != nil {
		return
	}
	if err = s.FeedQ.Push(feed.ID); err != nil {
		return
	}
	out = &feed.ID
	return
}
