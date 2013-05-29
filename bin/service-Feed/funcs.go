package main

import (
	"errors"
	"git.300brand.com/coverage"
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

	if err = s.Writer.SendOnce(nil, "UpdateFeed", in, out); err != nil {
		s.Log.Error(err.Error())
	}
	return
}
