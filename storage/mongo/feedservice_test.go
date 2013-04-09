package mongo

import (
	"git.300brand.com/coverage"
	"net/url"
	"testing"
)

func TestFeedIdMap(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	fs := NewFeedService(m)

	urlStr := []string{
		"http://www.google.com/feed1.rss",
		"http://www.google.com/feed2.rss",
		"http://www.google.com/feed3.rss",
		"http://www.google.com/feed4.rss",
		"http://www.google.com/feed5.rss",
		"http://www.google.com/feed6.rss",
		"http://www.google.com/feed7.rss",
		"http://www.google.com/feed8.rss",
		"http://www.google.com/feed9.rss",
		"http://www.google.com/feed10.rss",
	}

	urls := make([]*url.URL, len(urlStr))
	for i := range urlStr {
		urls[i], _ = url.Parse(urlStr[i])

		f := coverage.NewFeed()
		f.URL = urls[i]
		if err := fs.Update(f); err != nil {
			t.Fatal(err)
		}
	}

	ids, err := m.FeedIds(urls)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ids); l != len(urls) {
		t.Error("Length mismatch: %d", ids)
	}
}
