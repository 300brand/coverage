package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
)

// Filter funcs return true when the *html.Node passed in matches some
// condition.
type Filter func(n *html.Node) bool

func BlockElement(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	switch n.DataAtom {
	case atom.Article:
	case atom.Div:
	case atom.Footer:
	case atom.Header:
	case atom.Nav:
	case atom.Section:
	case atom.Td:
	default:
		return false
	}
	return true
}

func Comment(n *html.Node) bool {
	return n.Type == html.CommentNode
}

func Empty(n *html.Node) bool {
	return n.Type == html.ElementNode && n.FirstChild == nil
}

func Head(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Head
}

// Note this always returns false. This filter silently modifies the incoming
// *html.Node to change the tag to a div for future processing.
func NormalizeBlock(n *html.Node) bool {
	if BlockElement(n) {
		n.DataAtom = atom.Div
		n.Data = "div"
	}
	return false
}

func Script(n *html.Node) bool {
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
