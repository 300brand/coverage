package main

import (
	"git.300brand.com/coverage/skytypes"
	"net/http"
)

type Queue struct{}

func init() {
	s.RegisterService(new(Queue), "")
}

func (q *Queue) AddFeed(r *http.Request, in *skytypes.NullType, out *skytypes.ObjectId) (err error) {
	err = GetService("Queue").Send(nil, "AddFeed", *in, *out)
	return
}
