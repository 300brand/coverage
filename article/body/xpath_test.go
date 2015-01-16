package body

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/300brand/coverage"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type Set struct {
	File   string
	Expect string
	XPaths []string
}

var SamplesDir = "../samples"

func TestXPathP(t *testing.T) {
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

func TestXPathBrBr(t *testing.T) {
	b := []byte(`
<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<article>
			<header>
				<h1>Article Title</h1>
			</header>
			Article <a href="#">body</a>.<br><br
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
	// t.Skip("Missing body data")
	// return

	f, err := os.Open(filepath.Join(SamplesDir, "body-xpaths.json"))
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
			t.Errorf("[%d] [%s] %s", i, s.File, err)
		}
	}
}

func (s Set) Test(t *testing.T) (err error) {
	in, err := ioutil.ReadFile(filepath.Join(SamplesDir, s.File))
	if err != nil {
		return
	}
	expect, err := ioutil.ReadFile(filepath.Join(SamplesDir, s.Expect))
	if err != nil {
		return
	}
	body := new(coverage.Body)
	if err = XPath(in, s.XPaths, body); err != nil {
		return
	}
	if !bytes.Equal(body.Text, expect) {
		t.Errorf("Expect\n\n%s", expect)
		t.Errorf("Got\n\n%s", body.Text)
		t.Errorf("HTML\n\n%s", body.HTML)
		return fmt.Errorf("Did not get expected result")
	}
	return
}
