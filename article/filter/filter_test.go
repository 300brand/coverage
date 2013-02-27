package filter

import (
	"bytes"
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
		html.CommentNode:  false,
		html.DoctypeNode:  false,
		html.DocumentNode: false,
		html.ElementNode:  true,
		html.ErrorNode:    false,
		html.TextNode:     false,
	}
	for i, isBlock := range types {
		n.Type = i
		if BlockElement(n) != isBlock {
			t.Errorf("Expected %v for %d", isBlock, i)
		}
	}
}

func TestCommentType(t *testing.T) {
	n := &html.Node{}
	types := map[html.NodeType]bool{
		html.CommentNode:  true,
		html.DoctypeNode:  false,
		html.DocumentNode: false,
		html.ElementNode:  false,
		html.ErrorNode:    false,
		html.TextNode:     false,
	}
	for i, isBlock := range types {
		n.Type = i
		if Comment(n) != isBlock {
			t.Errorf("Expected %v for %d", isBlock, i)
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := map[string]bool{
		"<br>": false,
	}
	for test, empty := range tests {
		r := bytes.NewReader([]byte(test))
		doc, err := html.Parse(r)
		if err != nil {
			t.Error(err)
		}
		if Empty(doc) != empty {
			t.Errorf("Expected %v for `%s'", empty, test)
		}
	}
}
