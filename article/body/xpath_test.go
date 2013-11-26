package body

import (
	"encoding/json"
	"github.com/300brand/coverage"
	"os"
	"testing"
)

type Set struct {
	File   string
	Expect string
	XPaths []string
}

func TestXPath(t *testing.T) {
	b := []byte(`
<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<article>
			<header>
				<h1>Article Title</h1>
			</header>
			<p class="headline">Article <a href="#">body</a>.</p>
			More article body.<br><br>Third line.
		</article>
	</body>
</html>`)
	search := []string{
		"-//header",
		"//article",
	}
	body := new(coverage.Body)
	if err := XPath(b, search, body); err != nil {
		t.Fatal(err)
	}
	t.Logf("Body Text: %s", body.Text)
	t.Logf("Body HTML: %s", body.HTML)
}

func TestXPathSamples(t *testing.T) {
	f, err := os.Open("../samples/body-xpaths.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	v := make([]Set, 0, 128)
	if err := dec.Decode(&v); err != nil {
		t.Fatal(err)
	}
	for i, s := range v {
		if err := s.Test(t); err != nil {
			t.Errorf("[%d] %s", i, err)
		}
	}
}

func (s Set) Test(t *testing.T) (err error) {
	return
}
