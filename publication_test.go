package coverage

import (
	"net/url"
	"testing"
)

func TestPubFeedAdd(t *testing.T) {
	p := NewPublication()
	f := NewFeed()
	f.URL, _ = url.Parse("http://google.com")
	p.AddFeed(f)
	if len(p.Feeds) != 1 {
		t.Error("Invalid feed count: %d", len(p.Feeds))
	}
}

func TestPubDupeFeedID(t *testing.T) {
	p := NewPublication()
	f := NewFeed()

	f.URL, _ = url.Parse("http://google.com")
	p.AddFeed(f)

	f2 := NewFeed()
	f2.ID = f.ID
	f2.URL, _ = url.Parse("http://duckduckgo.com")

	if err := p.AddFeed(f2); err == nil {
		t.Error("Expected error")
	}
}
