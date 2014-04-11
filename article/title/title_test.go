package title

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var basicHTML = []byte(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>My Test Page</title>
	</head>
	<body>
		<article>
			<header>
				<h1>My Sample Article</h1>
				<div rel="title">Joe title</div>
			</header>
			<p>Some content would go here</p>
			<p>Some more content would show up here, too</p>
		</article>
		<article>
			<header>
				<h1>My Sample Article</h1>
				<div rel="by-title">By:  Bob The  title</div>
			</header>
			<p>Some content would go here</p>
			<p>Some more content would show up here, too</p>
		</article>
	</body>
	</html>
`)

func TestSuperBasic(t *testing.T) {
	xpaths := []string{
		`//*[@rel="title"]`,
	}
	title, err := Search(basicHTML, xpaths)
	if err != nil {
		t.Fatal(err)
	}
	if exp := "Joe title"; title != exp {
		t.Errorf("Found: %s; Expected: %s", title, exp)
	}
}

func TestBy(t *testing.T) {
	xpaths := []string{
		`substring-after(//*[@rel="by-title"], "By:")`,
	}
	title, err := Search(basicHTML, xpaths)
	if err != nil {
		t.Fatal(err)
	}
	if exp := "Bob The title"; title != exp {
		t.Errorf("Found: %s; Expected: %s", title, exp)
	}
}

func TestTitlesJSON(t *testing.T) {
	data, err := ioutil.ReadFile("../samples/titles.json")
	if err != nil {
		t.Fatal(err)
	}

	v := []struct {
		File   string
		XPaths []string
		Expect string
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatal(err)
	}

	for _, f := range v {
		html, err := ioutil.ReadFile("../samples/" + f.File)
		if err != nil {
			t.Fatal(err)
		}
		title, err := Search(html, f.XPaths)
		if err != nil {
			t.Error(err)
			continue
		}
		if title != f.Expect {
			t.Errorf("[%s] Got %s; Expected %s", f.File, title, f.Expect)
		}
	}
}
