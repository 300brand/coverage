package body

import (
	"errors"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
)

type Service struct{}

var _ service.Service = Service{}

func NewService() Service {
	return Service{}
}

func (s Service) Update(a *coverage.Article) (err error) {
	if a.HTML == nil {
		return errors.New("HTML not set, did you run the downloader service?")
	}
	if a.Body, err = GetBody(a.HTML); err != nil {
		return
	}
	a.Modified()
	return
}
