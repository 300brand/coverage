package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"testing"
)

func TestBlockElementDataAtom(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
	}
	atoms := map[atom.Atom]bool{
		atom.A:       false,
		atom.Article: true,
		atom.Body:    false,
		atom.Div:     true,
		atom.Footer:  true,
		atom.Head:    false,
		atom.Header:  true,
		atom.Input:   false,
		atom.Nav:     true,
		atom.Section: true,
		atom.Span:    false,
		atom.Table:   false,
		atom.Td:      true,
		atom.Tr:      false,
	}
	for i, isBlock := range atoms {
		n.DataAtom = i
		if BlockElement(n) != isBlock {
			t.Errorf("Expected %v for %s", !isBlock, i.String())
		}
	}
}

func TestBlockElementType(t *testing.T) {
	n := &html.Node{
		DataAtom: atom.Div,
	}
	types := map[html.NodeType]bool{
		html.ErrorNode:    false,
		html.TextNode:     false,
		html.DocumentNode: false,
		html.ElementNode:  true,
		html.CommentNode:  false,
		html.DoctypeNode:  false,
	}
	for i, isBlock := range types {
		n.Type = i
		if BlockElement(n) != isBlock {
			t.Errorf("Expected %v for %d", !isBlock, i)
		}
	}
}
