package main

import (
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/feed"
	"github.com/skynetservices/skynet"
)

func (s *FeedProcess) Process(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	if err = feed.Process(in); err != nil {
		s.Log.Error(err.Error())
		return
	}
	*out = *in
	s.Log.Debug(fmt.Sprintf("%v", out.URLs))
	return
}
