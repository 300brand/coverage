package downloader

import (
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
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
	return Article(a)
}

func Article(a *coverage.Article) error {
	a.Log.Service("downloader.ArticleService")
	r, err := Fetch(a.URL.String())
	if err != nil {
		return a.Log.Error(err)
	}
	a.LastCheck = time.Now()
	a.Text.HTML = r.Body
	if a.URL.String() != r.RealURL {
		a.Log.Debug("Updating URL from [%s] to [%s]", a.URL.String(), r.RealURL)
		if a.URL, err = url.Parse(r.RealURL); err != nil {
			return a.Log.Error(err)
		}
	}
	a.Modified("HTML")
	return nil
}
