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
	return SetBody(a)
}

func SetBody(a *coverage.Article) (err error) {
	if a.Text.HTML == nil {
		return a.Log.Error(errors.New("HTML not set, did you run the downloader service?"))
	}
	cleaned, err := CleanHTML(a.Text.HTML)
	if err != nil {
		return a.Log.Error(err)
	}
	if a.Text.Body, err = GetBody(cleaned); err != nil {
		return a.Log.Error(err)
	}
	a.Modified("Body")
	return
}
