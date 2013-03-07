package downloader

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"net/url"
	"time"
)

type Service struct {
	URL string
}

var _ service.Service = &Service{}

func NewService(url string) *Service {
	return &Service{URL: url}
}

func (s *Service) Update(a *coverage.Article) error {
	r, err := Fetch(s.URL)
	if err != nil {
		return err
	}
	a.Times.LastCheck = time.Now()
	a.HTML = r.Body
	if a.URL, err = url.Parse(r.RealURL); err != nil {
		return err
	}
	a.Modified()
	return nil
}
