package main

import (
	"errors"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/doozer/idqueue"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
)

func (s *Queue) AddFeed(ri *skynet.RequestInfo, in *skytypes.NullType, out *skytypes.ObjectId) (err error) {
	s.Log.Debug("Getting current queue")
	ids, err := s.FeedQ.Get()
	if err != nil && err != idqueue.ErrEOQ {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Debug(fmt.Sprintf("Queue retrieved with length of: %d", len(ids)))
	feed := &coverage.Feed{}
	skyIds := &skytypes.ObjectIds{Ids: ids}
	s.Log.Debug("Fetching oldest feed")
	if err = c.GetService("StorageReader", "", "", "").Send(nil, "OldestFeed", skyIds, feed); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Debug(fmt.Sprintf("Oldest feed ID: %s", feed.ID))
	if feed.ID.Hex() == "" {
		err = errors.New("No feed found")
		s.Log.Error(err.Error())
		return
	}
	if err = s.FeedQ.Push(feed.ID); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Debug("Added feed to queue")
	out.Id = feed.ID
	return
}
