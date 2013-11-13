package author

import (
	"testing"
)

func TestSuperBasic(t *testing.T) {
	html := []byte(`
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
		</body>
		</html>
	`)
	xpaths := []string{
		`//*[@rel="author"]`,
	}
	author, err := Search(html, xpaths)
	if err != nil {
		t.Fatal(err)
	}
	if exp := "Joe Author"; author != exp {
		t.Errorf("Found: %s; Expected: %s", author, exp)
	}
}
