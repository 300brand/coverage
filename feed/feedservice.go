package feed

import (
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/feed/normalizer"
	"github.com/300brand/coverage/feed/parser"
	"github.com/300brand/coverage/service"

	"fmt"
	_ "github.com/300brand/coverage/feed/parser/atom"
	_ "github.com/300brand/coverage/feed/parser/rdf"
	_ "github.com/300brand/coverage/feed/parser/rss"
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
		// TODO Using article.URL here causes the URL to change to the final,
		// cleaned version - need to either use strings or dereference into the
		// array in AddURL
		if !f.AddURL(article.URL) {
			continue
		}

		a := coverage.NewArticle()
		a.FeedId = f.ID
		a.PublicationId = f.PublicationId
		a.Title = article.Title
		a.URL = article.URL
		a.Published = article.Published
		f.Articles = append(f.Articles, a)
		f.Log.Debug("New article: %s", a.ID.Hex())
	}
	return nil
}
