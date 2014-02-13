package downloader

import (
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
)

type FeedService struct{}

var _ service.FeedService = FeedService{}

func NewFeedService() FeedService {
	return FeedService{}
}

func (s FeedService) Update(f *coverage.Feed) error {
	f.Log.Service("downloader.FeedService")
	return Feed(f)
}

func Feed(f *coverage.Feed) error {
	defer f.Downloaded()

	r, err := Fetch(f.URL)
	if err != nil {
		return f.Log.Error(err)
	}
	f.Content = r.Body
	return nil
}
