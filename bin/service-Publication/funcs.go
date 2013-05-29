package main

import (
	"errors"
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
	"strings"
)

func (s *Publication) Add(ri *skynet.RequestInfo, in *coverage.Publication, out *coverage.Publication) (err error) {
	// TODO Add check to make sure Publication with URL doesn't already exist
	errs := make([]string, 0, 2)
	if in.Title == "" {
		errs = append(errs, "Publication cannot have a blank title")
	}
	if in.URL == nil {
		errs = append(errs, "Publication cannot have a blank URL")
	}
	if len(errs) > 0 {
		errMsg := strings.Join(errs, "; ")
		s.Log.Error(errMsg)
		return errors.New(errMsg)
	}

	if err = s.Writer.SendOnce(nil, "SavePublication", in, out); err != nil {
		s.Log.Error(err.Error())
	}
	return
}
