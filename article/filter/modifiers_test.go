package filter

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"testing"
)

func TestAddBlockBreaks(t *testing.T) {
	test := `<!DOCTYPE html><html><head></head><body><br/>
<div></div>
<h1></h1>
<h2></h2>
<h3></h3>
<h4></h4>
<h5></h5>
<h6></h6>
<p></p>
</body></html>`

	r := bytes.NewReader(bytes.Replace([]byte(test), []byte{'\n'}, []byte{}, -1))
	n, err := html.Parse(r)
	if err != nil {
		t.Error(err)
	}

	// First node inside of body tag
	c := n.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild

	for c != nil {
		next := c.NextSibling
		AddBlockBreaks(c)
		c = next
	}

	out := &bytes.Buffer{}
	html.Render(out, n)
	if out.String() != test {
		t.Error("Replacement failure")
		t.Logf("Expected:\n%s", test)
		t.Logf("Got:\n%s", out.String())
	}
}

func TestDespace(t *testing.T) {
	t.Error("TODO")
}

func TestNormalizeBlock(t *testing.T) {
	t.Error("TODO")
}
