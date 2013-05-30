package main

import (
	"labix.org/v2/mgo/bson"
	"net/http"
)

type Queue struct{}

type QueueAddIn struct {
	Count int
}

type QueueAddOut struct {
	Added []bson.ObjectId
}

func init() {
	s.RegisterService(new(Queue), "")
}

func (q *Queue) AddFeed(r *http.Request, in *QueueAddIn, out *QueueAddOut) (err error) {
	out.Added = make([]bson.ObjectId, 0, in.Count)
	var feed struct{ Id bson.ObjectId }
	err = GetService("Queue").Send(nil, "AddFeed", Null, &feed)
	out.Added = append(out.Added, feed.Id)
	return
}
