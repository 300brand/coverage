package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
)

type ArticleService struct {
	m *Mongo
}

var _ service.ArticleService = &ArticleService{}

func NewArticleService(m *Mongo) *ArticleService {
	return &ArticleService{m: m}
}

func (s *ArticleService) Update(a *coverage.Article) error {
	return s.m.UpdateArticle(a)
}
