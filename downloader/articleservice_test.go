package downloader

import (
	"git.300brand.com/coverage"
	"net/url"
	"os"
	"testing"
)

func TestArticleService(t *testing.T) {
	s := NewArticleService()
	a := coverage.NewArticle()
	cwd, _ := os.Getwd()
	a.URL, _ = url.Parse("file://" + cwd + "/sample-download.txt")
	s.Update(a)
	if len(a.Text.HTML) != 447 {
		t.Errorf("URL not downloaded properly - %d bytes", len(a.Text.HTML))
	}
}
