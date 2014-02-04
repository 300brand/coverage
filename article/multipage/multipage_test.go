package multipage

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var simpleHTML = []byte(`<!DOCTYPE html>
<html>
<head></head>
<body>
	<a href="./">Prev</a>
	<a href="./1">1</a>
	<a href="2">2</a>
	<a href="?3">3</a>
	<!-- page 4 shouldn't register -->
	<a href="#4">4</a>
	<a href="./2">Next</a>
</body>
</html>
`)

func TestFind(t *testing.T) {
	links, err := FindLinks(simpleHTML)
	if err != nil {
		t.Fatalf("Error finding links: %s", err)
	}

	for _, l := range links {
		t.Logf("%v", l)
	}
}

// http://www.theregister.co.uk/2014/01/22/review_ouya_android_games_console/
// http://gcn.com/Articles/2013/12/03/big-data-tools-1.aspx?Page=1
// http://fcw.com/articles/2014/01/31/acquisition-games.aspx

func TestSamples(t *testing.T) {
	ts := httptest.NewServer(http.FileServer(http.Dir("samples")))
	defer ts.Close()

	d, err := os.Open("samples")
	if err != nil {
		t.Fatalf("Error opening samples dir: %s", err)
	}
	defer d.Close()

	fis, err := d.Readdir(0)
	if err != nil {
		t.Fatalf("Error reading samples dir: %s", err)
	}

	base, _ := url.Parse(ts.URL)
	for _, fi := range fis {
		if !fi.IsDir() {
			continue
		}
		subdir, _ := url.Parse(fi.Name())
		index := base.ResolveReference(subdir)
		resp, err := http.Get(index.String())
		if err != nil {
			t.Errorf("Error fetching %s: %s", index, err)
		}
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Problem reading %s: %s", index, err)
		}
		links, err := FindLinks(html)
		for _, link := range links {
			t.Logf("[%s] Link found: [%d] %s", index, link.Num, link.Url)
		}
	}
}
