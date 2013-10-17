package lexer

import (
	"errors"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
)

type ArticleService struct{}

var _ service.ArticleService = &ArticleService{}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) Update(a *coverage.Article) error {
	a.Log.Service("lexer.ArticleService")
	if len(a.Text.Body.Text) == 0 {
		return a.Log.Error(errors.New("Article body is empty, did you run body.ArticleService?"))
	}
	a.Text.Words.All = Words(a.Text.Body.Text)
	a.Text.Words.Keywords = Keywords(a.Text.Body.Text)
	a.Modified("Words")
	return nil
}
