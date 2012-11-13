package download

import (
	"git.300brand.com/coverage/parser"
	_ "git.300brand.com/coverage/parser/atom"
	"testing"
	"time"
)

func TestDownloadFeed(t *testing.T) {
	req := &FeedRequest{
		Timeout: time.Minute * 5,
		URL:     "http://www.theregister.co.uk/headlines.atom",
	}
	if _, err := downloadFeed(req); err != nil {
		t.Error(err)
	}
}

func TestDownloadFail(t *testing.T) {
	req := &FeedRequest{
		Timeout: time.Millisecond,
		URL:     "http://www.theregister.co.uk/headlines.atom",
	}
	if _, err := downloadFeed(req); err == nil {
		t.Error("Timeout expected")
	}
}

func TestFeedExtraction(t *testing.T) {
	req := &FeedRequest{
		Timeout: time.Minute,
		URL:     "http://www.theregister.co.uk/headlines.atom",
	}
	resp, err := downloadFeed(req)
	if err != nil {
		t.Error(err)
	}
	feed, err := parser.Parse(resp.Body, req.Type)
	if err != nil {
		t.Error(err)
	}
	if len(feed.Articles) != 50 {
		t.Errorf("Did not get expected Article count, %d", len(feed.Articles))
	}
}
