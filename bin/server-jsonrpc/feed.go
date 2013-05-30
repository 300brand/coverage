package main

import (
	"git.300brand.com/coverage"
	"net/http"
	"net/url"
)

type Feed struct{}

type FeedAddArgs struct {
	RemoteId uint64
	URL      string
}

func init() {
	s.RegisterService(new(Feed), "")
}

func (f *Feed) Add(r *http.Request, in *FeedAddArgs, out *coverage.Feed) (err error) {
	feedIn := coverage.NewFeed()
	feedIn.ObjectId = in.RemoteId
	if feedIn.URL, err = url.Parse(in.URL); err != nil {
		return
	}
	err = GetService("Feed").Send(nil, "Add", feedIn, out)
	return
}
