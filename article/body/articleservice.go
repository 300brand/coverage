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
	if a.HTML == nil {
		return errors.New("HTML not set, did you run the downloader service?")
	}
	if a.Body, err = GetBody(a.HTML); err != nil {
		return
	}
	a.Modified()
	return
}
