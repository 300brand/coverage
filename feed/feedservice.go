package feed

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/feed/normalizer"
	"git.300brand.com/coverage/feed/parser"
	"git.300brand.com/coverage/service"

	"fmt"
	_ "git.300brand.com/coverage/feed/parser/atom"
	_ "git.300brand.com/coverage/feed/parser/rdf"
	_ "git.300brand.com/coverage/feed/parser/rss"
)

type FeedService struct{}

var _ service.FeedService = &FeedService{}

func NewFeedService() *FeedService {
	return &FeedService{}
}

func (s *FeedService) Update(f *coverage.Feed) error {
	f.Log.Service("feed.FeedService")
	return Process(f)
}

func Process(f *coverage.Feed) error {
	// Parse
	d, err := parser.Parse(f.Content)
	if err != nil {
		return f.Log.Error(fmt.Errorf("Decoder error: %s", err))
	}

	// Normalize
	n := &normalizer.Default{}
	if err := n.Normalize(d); err != nil {
		return f.Log.Error(fmt.Errorf("Normalizer error: %s", err))
	}

	// Apply and let dry
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
