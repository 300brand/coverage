package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/doozer/idqueue"
	"github.com/skynetservices/skynet"
	"labix.org/v2/mgo/bson"
)

type AddFeedResponse struct {
	Id bson.ObjectId
}

func (s *Queue) AddFeed(ri *skynet.RequestInfo, in *interface{}, out *AddFeedResponse) (err error) {
	ids, err := s.FeedQ.Get()
	if err != nil && err != idqueue.ErrEOQ {
		return
	}
	feed := &coverage.Feed{}
	if err = c.GetService("StorageReader", "", "", "").Send(ri, "OldestFeed", ids, feed); err != nil {
		return
	}
	if err = s.FeedQ.Push(feed.ID); err != nil {
		return
	}
	out.Id = feed.ID
	return
}
