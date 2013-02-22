package body

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"git.300brand.com/coverage/page/filter"
)

var cleanFilters = filter.Filters{
	filter.Head,
	filter.Style,
	filter.Script,
	filter.Comment,
	filter.NormalizeBlock,
}

func CleanHTML(b []byte) (cleaned []byte, err error) {
	r := bytes.NewReader(b)
	doc, err := html.Parse(r)
	if err != nil {
		return
	}
	cleanDOM(doc)
	buf := &bytes.Buffer{}
	html.Render(buf, doc)
	cleaned = buf.Bytes()
	return
}

func cleanDOM(n *html.Node) {
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		if cleanFilters.Any(c) {
			n.RemoveChild(c)
		} else {
			cleanDOM(c)
		}
		c = next
	}
}
