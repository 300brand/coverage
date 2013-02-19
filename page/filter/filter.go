package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
)

type Filter func(n *html.Node) bool

func AnyValid(n *html.Node, filters []Filter) bool {
	for _, f := range filters {
		if f(n) {
			return true
		}
	}
	return false
}

func Comments(n *html.Node) bool {
	return n.Type == html.CommentNode
}

func Empty(n *html.Node) bool {
	return n.Type == html.ElementNode && n.FirstChild == nil
}

func Head(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Head
}

func Scripts(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	switch n.DataAtom {
	case atom.Script:
	case atom.Noscript:
	default:
		return false
	}
	return true
}

func Style(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Style
}

func Text(n *html.Node) bool {
	return n.Type == html.TextNode
}
