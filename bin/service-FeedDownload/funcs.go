package main

import (
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/downloader"
	"github.com/skynetservices/skynet"
)

func (s *FeedDownload) Download(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	if err = downloader.Feed(in); err != nil {
		s.Log.Error(err.Error())
		return
	}
	*out = *in
	s.Log.Debug(fmt.Sprintf("RSS XML len: %d", len(out.Content)))
	return
}
