package main

import (
	"errors"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
	"strings"
)

func (s *Feed) Add(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	// TODO Add check to make sure Feed with URL doesn't already exist
	errs := make([]string, 0, 1)
	if in.URL == nil {
		errs = append(errs, "Feed cannot have a blank URL")
	}
	if len(errs) > 0 {
		errMsg := strings.Join(errs, "; ")
		s.Log.Error(errMsg)
		return errors.New(errMsg)
	}

	if err = s.StorageWriter.SendOnce(nil, "SaveFeed", in, out); err != nil {
		s.Log.Error(err.Error())
	}
	return
}

func (s *Feed) Get(ri *skynet.RequestInfo, in *skytypes.ObjectId, out *coverage.Feed) (err error) {
	return s.StorageReader.Send(ri, "GetFeed", in, out)
}

func (s *Feed) Oldest(ri *skynet.RequestInfo, in *skytypes.ObjectIds, out *coverage.Feed) (err error) {
	return s.StorageReader.Send(ri, "OldestFeed", in, out)
}

func (s *Feed) Process(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	if err = s.FeedDownload.Send(ri, "Download", in, out); err != nil {
		return
	}
	*in = *out
	if err = s.FeedProcess.Send(ri, "Process", in, out); err != nil {
		return
	}
	*in = *out
	if err = s.StorageWriter.Send(ri, "SaveFeed", in, out); err != nil {
		return
	}
	// TODO Save all articles to database; Add ArticleIds to queue
	return
}
