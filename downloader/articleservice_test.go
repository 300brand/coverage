package downloader

import (
	"github.com/300brand/coverage"
	"net/url"
	"os"
	"testing"
)

func TestArticleService(t *testing.T) {
	a := coverage.NewArticle()
	cwd, _ := os.Getwd()
	a.URL, _ = url.Parse("file://" + cwd + "/sample-download.txt")
	if err := Article(a); err != nil {
		t.Fatal(err)
	}
	if len(a.Text.HTML) != 447 {
		t.Errorf("URL not downloaded properly - %d bytes", len(a.Text.HTML))
	}
}

func TestArticleMetaRefresh(t *testing.T) {
	a := coverage.NewArticle()
	cwd, _ := os.Getwd()
	a.URL, _ = url.Parse("file://" + cwd + "/sample-meta-refresh.html")
	if err := Article(a); err != nil {
		t.Fatal(err)
	}
	if len(a.Text.HTML) != 447 {
		t.Errorf("URL not downloaded properly - %d bytes", len(a.Text.HTML))
	}
}
