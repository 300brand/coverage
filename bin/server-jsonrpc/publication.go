package main

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"net/http"
	"net/url"
)

type Publication struct{}

type PubAddArgs struct {
	Title string
	URL   string
}

type PubAddReply struct {
	Id bson.ObjectId
}

func init() {
	s.RegisterService(new(Publication), "")
}

func (p *Publication) Add(r *http.Request, in *PubAddArgs, out *PubAddReply) (err error) {

	pubIn := coverage.NewPublication()
	pubIn.Title = in.Title
	if pubIn.URL, err = url.Parse(in.URL); err != nil {
		return
	}
	pubOut := &coverage.Publication{}

	err = GetService("Publication").Send(nil, "Add", pubIn, pubOut)

	out.Id = pubOut.ID
	return
}
