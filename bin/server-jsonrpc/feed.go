package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
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

func (f *Feed) Download(r *http.Request, in *skytypes.ObjectId, out *coverage.Feed) (err error) {
	if err = GetService("Feed").Send(nil, "Get", in, out); err != nil {
		return
	}
	if err = GetService("FeedDownload").Send(nil, "Download", out, out); err != nil {
		return
	}
	return
}

func (f *Feed) Get(r *http.Request, in *skytypes.ObjectId, out *coverage.Feed) (err error) {
	return GetService("Feed").Send(nil, "Get", in, out)
}

func (f *Feed) Oldest(r *http.Request, in *skytypes.NullType, out *coverage.Feed) (err error) {
	return GetService("Feed").Send(nil, "Oldest", in, out)
}

func (f *Feed) ProcessId(r *http.Request, in *skytypes.ObjectId, out *coverage.Feed) (err error) {
	feed := &coverage.Feed{}
	srv := GetService("Feed")
	if err = srv.Send(nil, "Get", in, feed); err != nil {
		return
	}
	return srv.Send(nil, "Process", feed, out)
}
