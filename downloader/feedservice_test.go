package downloader

import (
	"git.300brand.com/coverage"
	"net/url"
	"os"
	"testing"
)

func TestFeedService(t *testing.T) {
	s := NewFeedService()
	f := coverage.NewFeed()
	cwd, _ := os.Getwd()
	f.URL, _ = url.Parse("file://" + cwd + "/sample-download.txt")
	s.Update(f)
	if len(f.Content) != 447 {
		t.Errorf("URL not downloaded properly - %d bytes", len(f.Content))
	}
}
