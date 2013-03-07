package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
)

type ArticleService struct{}

var _ service.ArticleService = &ArticleService{}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) Update(a *coverage.Article) error {

	return nil
}
