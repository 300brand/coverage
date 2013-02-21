package body

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"git.300brand.com/coverage/page/filter"
	"os"
	//"github.com/moovweb/gokogiri"
)

var cleanFilters = []filter.Filter{
	filter.Head,
	filter.Style,
	filter.Script,
	filter.Comment,
	filter.NormalizeBlock,
}

func CleanDOM(n *html.Node) {
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		if filter.AnyValid(c, cleanFilters) {
			n.RemoveChild(c)
		} else {
			CleanDOM(c)
		}
		c = next
	}
}

func GetBody(b []byte) (body string, err error) {
	r := bytes.NewReader(b)
	doc, err := html.Parse(r)
	if err != nil {
		return
	}
	CleanDOM(doc)
	html.Render(os.Stdout, doc)
	//
	//, err := gokogiri.ParseHtml(b)
	//if err != nil {
	//	return
	//}
	return
}
