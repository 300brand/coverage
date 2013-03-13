package body

import (
	"errors"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
)

type ArticleService struct{}

var _ service.ArticleService = ArticleService{}

func NewArticleService() ArticleService {
	return ArticleService{}
}

func (s ArticleService) Update(a *coverage.Article) (err error) {
	a.Log.Service("body.ArticleService")
	if a.HTML == nil {
		return a.Log.Error(errors.New("HTML not set, did you run the downloader service?"))
	}
	cleaned, err := CleanHTML(a.HTML)
	if err != nil {
		return a.Log.Error(err)
	}
	if a.Body, err = GetBody(cleaned); err != nil {
		return a.Log.Error(err)
	}
	a.Modified()
	return
}
