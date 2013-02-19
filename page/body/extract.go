package body

import (
	"code.google.com/p/go.net/html"
	"git.300brand.com/coverage/page/filter"
)

var cleanFilters = []filter.Filter{
	filter.Head,
	filter.Style,
	filter.Scripts,
	filter.Comments,
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
