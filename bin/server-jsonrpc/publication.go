package main

import (
	"git.300brand.com/coverage"
	"net/http"
	"net/url"
)

type Publication struct{}

type PubAddArgs struct {
	Title string
	URL   string
}

func init() {
	s.RegisterService(new(Publication), "")
}

func (p *Publication) Add(r *http.Request, in *PubAddArgs, out *coverage.Publication) (err error) {

	pubIn := coverage.NewPublication()
	pubIn.Title = in.Title
	if pubIn.URL, err = url.Parse(in.URL); err != nil {
		return
	}
	err = GetService("Publication").Send(nil, "Add", pubIn, out)
	return
}
