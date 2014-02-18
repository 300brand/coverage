package multipage

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestSamples(t *testing.T) {
	tests := []struct {
		Url    *url.URL
		Expect []int
	}{
		{&url.URL{Path: "fcw-games/index.html"}, []int{2}},
		{&url.URL{Path: "gcn-big-data/index.html"}, []int{2, 3}},
		{&url.URL{Path: "theregister-ouya/index.html"}, []int{2}},
	}

	ts := httptest.NewServer(http.FileServer(http.Dir("samples")))
	defer ts.Close()

	base, _ := url.Parse(ts.URL)
	for _, test := range tests {
		index := base.ResolveReference(test.Url)
		resp, err := http.Get(index.String())
		if err != nil {
			t.Errorf("Error fetching %s: %s", index, err)
		}
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Problem reading %s: %s", index, err)
		}
		pages, err := Pages(index, html)
		if err != nil {
			t.Errorf("Error finding pages: %s", err)
		}
		if lP, lE := len(pages), len(test.Expect); lP != lE {
			t.Errorf("Invalid number of pages found: %d != %d", lP, lE)
			continue
		}
		for i := range pages {
			if got, exp := pages[i].Num, test.Expect[i]; got != exp {
				t.Errorf("Got page %d; Expected %d", got, exp)
			}
		}
	}
}
