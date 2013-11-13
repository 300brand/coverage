package author

import (
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
				<div rel="author">Joe Author</div>
			</header>
			<p>Some content would go here</p>
			<p>Some more content would show up here, too</p>
		</article>
		<article>
			<header>
				<h1>My Sample Article</h1>
				<div rel="by-author">By:  Joe   Author</div>
			</header>
			<p>Some content would go here</p>
			<p>Some more content would show up here, too</p>
		</article>
	</body>
	</html>
`)

func TestSuperBasic(t *testing.T) {
	xpaths := []string{
		`//*[@rel="author"]`,
	}
	author, err := Search(basicHTML, xpaths)
	if err != nil {
		t.Fatal(err)
	}
	if exp := "Joe Author"; author != exp {
		t.Errorf("Found: %s; Expected: %s", author, exp)
	}
}

func TestBy(t *testing.T) {
	xpaths := []string{
		`substring-after(//*[@rel="by-author"], "By:")`,
	}
	author, err := Search(basicHTML, xpaths)
	if err != nil {
		t.Fatal(err)
	}
	if exp := "Joe Author"; author != exp {
		t.Errorf("Found: %s; Expected: %s", author, exp)
	}
}
