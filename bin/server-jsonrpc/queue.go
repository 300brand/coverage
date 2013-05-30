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
	out = &QueueAddOut{
		Added: make([]bson.ObjectId, 0, in.Count),
	}
	srv := c.GetService("Queue", "", "", "")
	err = srv.Send(nil, "AddFeed", nil, &out.Added)
	return
}
