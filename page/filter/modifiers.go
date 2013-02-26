// These always return false. These filters silently modify the incoming
// *html.Node to add or change the DOM
package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"strings"
)

// Add newlines after block-level elements
func AddBlockBreaks(n *html.Node) bool {
	switch n.DataAtom {
	case atom.Br:
	case atom.Div:
	case atom.H1:
	case atom.H2:
	case atom.H3:
	case atom.H4:
	case atom.H5:
	case atom.H6:
	case atom.P:
	default:
		return false
	}
	nl := &html.Node{
		Type: html.TextNode,
		Data: "\n",
	}
	if n.NextSibling != nil {
		n.Parent.InsertBefore(nl, n.NextSibling)
	} else {
		n.Parent.AppendChild(nl)
	}
	return false
}

// Converts all internal content spacing into a single space
func Despace(n *html.Node) bool {
	if n.Type != html.TextNode {
		return false
	}
	n.Data = strings.Trim(n.Data, " \t\r\n")
	n.Data = despaceRegex.ReplaceAllString(n.Data, " ")
	return false
}

// Converts all block-level tags to div tags
func NormalizeBlock(n *html.Node) bool {
	if BlockElement(n) {
		n.DataAtom = atom.Div
		n.Data = "div"
	}
	return false
}
