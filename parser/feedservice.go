package parser

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/parser/normalizer"
	"git.300brand.com/coverage/service"

	_ "git.300brand.com/coverage/parser/atom"
	_ "git.300brand.com/coverage/parser/rdf"
	_ "git.300brand.com/coverage/parser/rss"
)

type FeedService struct{}

var _ service.FeedService = &FeedService{}

func (s *FeedService) Update(f *coverage.Feed) error {
	n := &normalizer.Default{}
	if err := Normalize(f.Content, n); err != nil {
		return f.Log.Error(err)
	}
	for _, article := range n.Articles {
		a := coverage.NewArticle()
		a.Title = article.Title
		a.URL = &article.URL
		a.Published = article.Published
		f.Articles = append(f.Articles, a)
	}
	return nil
}
