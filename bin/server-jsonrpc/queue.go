package main

import (
	"labix.org/v2/mgo/bson"
	"net/http"
)

type Queue struct{}

func init() {
	s.RegisterService(new(Queue), "")
}

type QueueAddIn struct {
	Count uint
}

type QueueAddOut struct {
	Added []bson.ObjectId
}

func (q *Queue) AddFeed(r *http.Request, in *QueueAddIn, out *QueueAddOut) (err error) {
	out.Added = make([]bson.ObjectId, 0, in.Count)
	var id bson.ObjectId
	err = GetService("Queue").Send(nil, "AddFeeds", nil, &id)
	out.Added = append(out.Added, id)
	return
}
