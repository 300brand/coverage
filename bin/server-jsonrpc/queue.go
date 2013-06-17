package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"net/http"
)

type Queue struct{}

func init() {
	s.RegisterService(new(Queue), "")
}

func (q *Queue) AddFeed(r *http.Request, in *skytypes.NullType, out *skytypes.ObjectId) (err error) {
	return GetService("Queue").Send(nil, "AddFeed", in, out)
}

func (q *Queue) AddOneFeed(r *http.Request, in *skytypes.NullType, out *skytypes.NullType) (err error) {
	cmd := &skytypes.ClockCommand{
		Command: "once",
	}
	return GetService("Manager").Send(nil, "QueueFeedAdder", cmd, skytypes.Null)
}

func (q *Queue) NextFeed(r *http.Request, in *skytypes.NullType, out *coverage.Feed) (err error) {
	return GetService("Queue").Send(nil, "NextFeed", in, out)
}
