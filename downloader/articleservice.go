package downloader

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"net/url"
	"time"
)

type ArticleService struct {
	URL string
}

var _ service.ArticleService = ArticleService{}

func NewArticleService(url string) ArticleService {
	return ArticleService{URL: url}
}

func (s ArticleService) Update(a *coverage.Article) error {
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
