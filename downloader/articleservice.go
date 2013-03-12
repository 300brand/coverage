package downloader

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"net/url"
	"time"
)

type ArticleService struct {
}

var _ service.ArticleService = ArticleService{}

func NewArticleService() ArticleService {
	return ArticleService{}
}

func (s ArticleService) Update(a *coverage.Article) error {
	a.Log.Service("downloader.ArticleService")
	r, err := Fetch(a.URL.String())
	if err != nil {
		return a.Log.Error(err)
	}
	a.LastCheck = time.Now()
	a.HTML = r.Body
	if a.URL.String() != r.RealURL {
		a.Log.Debug("Updating URL from [%s] to [%s]", a.URL.String(), r.RealURL)
		if a.URL, err = url.Parse(r.RealURL); err != nil {
			return a.Log.Error(err)
		}
	}
	a.Modified()
	return nil
}
