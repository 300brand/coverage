package body

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"github.com/300brand/coverage/article/filter"
	"github.com/jbaikge/logger"
)

var cleanFilters = filter.Filters{
	filter.Aside,
	filter.Input,
	filter.Spaces,
	filter.Head,
	filter.Style,
	filter.Script,
	filter.Comment,
	filter.NormalizeBlock,
	filter.Despace,
	filter.Empty,
	filter.AddBlockBreaks,
	filter.TranslateUnicode,
}

func CleanHTML(b []byte) (cleaned []byte, err error) {
	logger.Trace.Print("CleanHTML: called")
	r := bytes.NewReader(b)
	doc, err := html.Parse(r)
	if err != nil {
		logger.Error.Printf("CleanHTML: %s", err)
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
