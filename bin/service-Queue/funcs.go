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
	if len(ids) >= s.FeedQ.Max {
		return idqueue.ErrFull
	}
	feed := &coverage.Feed{}
	skyIds := &skytypes.ObjectIds{Ids: ids}
	s.Log.Debug("Fetching oldest feed")
	if err = s.Feed.Send(nil, "Oldest", skyIds, feed); err != nil {
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

func (s *Queue) NextFeed(ri *skynet.RequestInfo, in *skytypes.NullType, out *coverage.Feed) (err error) {
	id := &skytypes.ObjectId{}
	id.Id, err = s.FeedQ.Unshift()
	if err != nil {
		return
	}
	s.Log.Debug(fmt.Sprintf("Got ID: %s", id.Id))
	return s.Feed.Send(nil, "Get", id, out)
}
