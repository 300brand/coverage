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
	a.URL, _ = url.Parse("file://" + cwd + "/../article/samples/AOLGov.html")
	s.Update(a)
	if len(a.HTML) != 60539 {
		t.Error("URL not downloaded properly")
	}
}
