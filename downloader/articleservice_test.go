package downloader

import (
	"github.com/300brand/coverage"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var ts *httptest.Server

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/article/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html><html><head></head><body><p>Main Article</p><a href="/article">1</a><a href="./2">2</a><a href="/article/3">3</a></body></html>`))
	})
	mux.HandleFunc("/article/2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html><html><head></head><body><p>Second Page</p><a href="/article">1</a><a href="./2">2</a><a href="/article/3">3</a></body></html>`))
	})
	mux.HandleFunc("/article/3", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html><html><head></head><body><p>Third Page</p><a href="/article">1</a><a href="./2">2</a><a href="/article/3">3</a></body></html>`))
	})
	ts = httptest.NewServer(mux)
}

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

func TestArticleMultipage(t *testing.T) {
	a := coverage.NewArticle()
	a.URL, _ = url.Parse(ts.URL + "/article")

	if err := Article(a); err != nil {
		t.Errorf("Error downloading: %s", err)
	}

	if l := len(a.Text.Pages); l != 2 {
		t.Errorf("Improper page count: %d; Expected %d", l, 2)
	}
}
