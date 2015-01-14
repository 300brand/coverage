package filter

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

// Filter funcs return true when the *html.Node passed in matches some
// condition.
type Filter func(n *html.Node) bool

func Aside(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Aside
}

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
	if n.Type != html.ElementNode {
		return false
	}
	c := n.FirstChild
	// Only child is a text node of spaces
	if c != nil && c == n.LastChild && Spaces(c) {
		return true
	}
	// Special-case inline tags
	switch n.DataAtom {
	case atom.Br:
	default:
		return c == nil
	}
	return false
}

func Input(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	switch n.DataAtom {
	case atom.Input:
	case atom.Select:
	case atom.Textarea:
	default:
		return false
	}
	return true
}

func Head(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Head
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

func Spaces(n *html.Node) bool {
	return Text(n) && strings.Trim(n.Data, " \t\r\n") == ""
}

func Style(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Style
}

func Text(n *html.Node) bool {
	return n.Type == html.TextNode
}
