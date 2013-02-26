package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"regexp"
	"strings"
)

// Filter funcs return true when the *html.Node passed in matches some
// condition.
type Filter func(n *html.Node) bool

var despaceRegex = regexp.MustCompile("(  +|[\t\r\n])+")

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
	if n.Type != html.ElementNode || n.FirstChild != nil {
		return false
	}
	switch n.DataAtom {
	case atom.Br:
	default:
		return true
	}
	return false
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
	return n.Type == html.TextNode && strings.Trim(n.Data, " \t\r\n") == ""
}

func Style(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.Style
}

func Text(n *html.Node) bool {
	return n.Type == html.TextNode
}
