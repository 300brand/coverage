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

func NewFeedService() *FeedService {
	return &FeedService{}
}

func (s *FeedService) Update(f *coverage.Feed) error {
	f.Log.Service("parser.FeedService")

	n := &normalizer.Default{}
	if err := Normalize(f.Content, n); err != nil {
		return f.Log.Error(err)
	}
	for _, article := range n.Articles {
		if !f.AddURL(article.URL) {
			continue
		}

		a := coverage.NewArticle()
		a.FeedId = f.ID
		a.Title = article.Title
		a.URL = article.URL
		a.Published = article.Published
		f.Articles = append(f.Articles, a)
		f.Log.Debug("New article: %s", a.ID.Hex())
	}
	return nil
}
