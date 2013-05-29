package main

import (
	"git.300brand.com/coverage"
	"net/http"
	"net/url"
)

type Feed struct{}

type FeedAddArgs struct {
	URL string
}

func init() {
	s.RegisterService(new(Feed), "")
}

func (p *Feed) Add(r *http.Request, in *FeedAddArgs, out *coverage.Feed) (err error) {
	feedIn := coverage.NewFeed()
	if feedIn.URL, err = url.Parse(in.URL); err != nil {
		return
	}
	err = GetService("Feed").Send(nil, "Add", feedIn, out)
	return
}