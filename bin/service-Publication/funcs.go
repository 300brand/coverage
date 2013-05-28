package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
)

func (s *Publication) Add(ri *skynet.RequestInfo, req *coverage.Publication, resp *coverage.Publication) (err error) {
	*resp = *req
	errs := make([]string, 0, 2)
	if req.Title == "" {
		errs = append(errs, "Publication cannot have a blank title")
	}
	if req.URL == nil {
		errs = append(errs, "Publication cannot have a blank URL")
	}
	if len(errs) > 0 {
		errMsg := strings.Join(errs, "; ")
		s.Log.Error(errMsg)
		return errors.New(errMsg)
	}

	if err = s.m.UpdatePublication(req); err != nil {
		s.Log.Error(err.Error())
	}
	return
}
