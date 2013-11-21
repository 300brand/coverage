package published

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
				<div rel="date">Thu Nov 21 11:47:16 2013</div>
			</header>
			<p>Some content would go here</p>
			<p>Some more content would show up here, too</p>
		</article>
		<article>
			<header>
				<h1>Dates From Sites Around the Net</h1>
			</header>
			<div rel="pcmag"> November 20, 2013</div>
			<div rel="idg">20.11.2013 kl 22:36 | IDG News Service\San Francisco Bureau</div>
		</article>
	</body>
	</html>
`)

func TestPublished(t *testing.T) {
	Layout := `Jan 02, 2006 15:04`
	tests := map[string]string{
		`//*[@rel="date"]~~Mon Jan _2 15:04:05 2006`:                    "Nov 21, 2013 11:47",
		`//*[@rel="pcmag"]~~January _2, 2006`:                           "Nov 20, 2013 00:00",
		`substring-before(//*[@rel="idg"], ' | ')~~02.01.2006 kl 15:04`: "Nov 20, 2013 22:36",
	}
	for test, expect := range tests {
		date, err := Search(basicHTML, []string{test})
		if err != nil {
			t.Fatal(err)
		}
		if got := date.Format(Layout); got != expect {
			t.Errorf("Incorrect date: %s; Expected: %s", got, expect)
		}
	}
}
