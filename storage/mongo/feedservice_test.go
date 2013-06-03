package mongo

import (
	"fmt"
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"net/url"
	"testing"
	"time"
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

func TestOldestOne(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	fs := NewFeedService(m)
	f := coverage.NewFeed()
	fs.Update(f)

	out, err := m.GetOldestFeed([]bson.ObjectId{})
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != f.ID {
		t.Error("Feed not found")
	}
}

func TestOldestMany(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	fs := NewFeedService(m)
	feeds := []*coverage.Feed{
		&coverage.Feed{
			ID:           bson.NewObjectId(),
			LastDownload: time.Now().Add(-5 * time.Minute),
		},
		&coverage.Feed{
			ID:           bson.NewObjectId(),
			LastDownload: time.Now().Add(-10 * time.Minute),
		},
		&coverage.Feed{
			ID:           bson.NewObjectId(),
			LastDownload: time.Now().Add(-7 * time.Minute),
		},
	}
	for i, f := range feeds {
		f.URL, _ = url.Parse(fmt.Sprintf("http://google.com/rss/%d", i))
		fs.Update(f)
	}

	order := []bson.ObjectId{feeds[1].ID, feeds[2].ID, feeds[0].ID}
	ignore := make([]bson.ObjectId, 0, len(feeds))
	for i := range feeds {
		f, err := m.GetOldestFeed(ignore)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s %s", f.ID, f.LastDownload)
		if f.ID != order[i] {
			t.Error("Incorrect order: Expected %s; Got %s", order[i], f.ID)
		}
		ignore = append(ignore, f.ID)
	}
}
