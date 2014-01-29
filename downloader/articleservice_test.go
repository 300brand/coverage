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

func TestArticleMetaRedirectDownload(t *testing.T) {
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

func TestArticleMetaRedirect(t *testing.T) {
	u, _ := url.Parse("http://google.com/test")
	body := []byte(`<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="refresh" content="0; passed">
	</head>
	<body></body>
</html>`)
	newUrl, err := metaRedirect(body, u)
	if err != nil {
		t.Fatal(err)
	}
	if newUrl.String() != "http://google.com/passed" {
		t.Fatal("Invalid redirect URL: %s", newUrl)
	}
}
