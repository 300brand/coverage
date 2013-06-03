package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/feed"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
	"labix.org/v2/mgo/bson"
)

func (s *FeedDownload) Process(ri *skynet.RequestInfo, in *coverage.Feed, out *skytypes.ObjectId) (err error) {
	out.Id = bson.NewObjectId()
	if err = downloader.Feed(in); err != nil {
		s.Log.Error(err.Error())
		return
	}
	// TODO save feed to DB
	if err = feed.Process(in); err != nil {
		s.Log.Error(err.Error())
		return
	}
	return
}
