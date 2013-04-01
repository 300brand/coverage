package lexer

import (
	"errors"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
)

type ArticleService struct{}

var _ service.ArticleService = &ArticleService{}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) Update(a *coverage.Article) error {
	a.Log.Service("lexer.ArticleService")
	if len(a.Body.Text) == 0 {
		return a.Log.Error(errors.New("Article body is empty, did you run body.ArticleService?"))
	}
	a.Words.All = Words(a.Body.Text)
	a.Words.Keywords = Keywords(a.Body.Text)
	a.Modified("Words")
	return nil
}
