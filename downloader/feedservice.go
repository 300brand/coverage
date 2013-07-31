package downloader

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
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

	r, err := Fetch(f.URL.String())
	if err != nil {
		return f.Log.Error(err)
	}
	f.Content = r.Body
	return nil
}
