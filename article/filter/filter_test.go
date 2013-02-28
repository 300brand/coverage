package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"strings"
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
			t.Errorf("Expected %v for %s", isBlock, i.String())
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
	for i, isType := range types {
		n.Type = i
		if BlockElement(n) != isType {
			t.Errorf("Expected %v for %d", isType, i)
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
	for i, isType := range types {
		n.Type = i
		if Comment(n) != isType {
			t.Errorf("Expected %v for %d", isType, i)
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := map[string]bool{
		"<br>":                   false,
		"<div></div>":            true,
		"<div><div></div></div>": false,
	}
	for test, empty := range tests {
		r := strings.NewReader(test)
		doc, err := html.Parse(r)
		if err != nil {
			t.Error(err)
		}
		// Path makes up for the fact that html.Parse() establishes
		// the entire HTML tree:
		// <html><head></head><body>...</body></html>
		node := doc.FirstChild.FirstChild.NextSibling.FirstChild
		if Empty(node) != empty {
			t.Errorf("Expected %v for `%s'", empty, test)
		}
	}
}

func TestHeadDataAtom(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
	}
	atoms := map[atom.Atom]bool{
		atom.Body:  false,
		atom.Div:   false,
		atom.Head:  true,
		atom.Html:  false,
		atom.Style: false,
	}
	for i, isHead := range atoms {
		n.DataAtom = i
		if Head(n) != isHead {
			t.Errorf("Expected %v for %s", isHead, i.String())
		}
	}
}

func TestHeadType(t *testing.T) {
	n := &html.Node{
		DataAtom: atom.Head,
	}
	types := map[html.NodeType]bool{
		html.CommentNode:  false,
		html.DoctypeNode:  false,
		html.DocumentNode: false,
		html.ElementNode:  true,
		html.ErrorNode:    false,
		html.TextNode:     false,
	}
	for i, isType := range types {
		n.Type = i
		if Head(n) != isType {
			t.Errorf("Expected %v for %d", isType, i)
		}
	}
}

func TestScriptDataAtom(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
	}
	atoms := map[atom.Atom]bool{
		atom.Body:     false,
		atom.Div:      false,
		atom.Html:     false,
		atom.Noscript: true,
		atom.Script:   true,
		atom.Style:    false,
	}
	for i, isScript := range atoms {
		n.DataAtom = i
		if Script(n) != isScript {
			t.Errorf("Expected %v for %s", isScript, i.String())
		}
	}
}

func TestScriptType(t *testing.T) {
	n := &html.Node{
		DataAtom: atom.Script,
	}
	types := map[html.NodeType]bool{
		html.CommentNode:  false,
		html.DoctypeNode:  false,
		html.DocumentNode: false,
		html.ElementNode:  true,
		html.ErrorNode:    false,
		html.TextNode:     false,
	}
	for i, isType := range types {
		n.Type = i
		if Script(n) != isType {
			t.Errorf("Expected %v for %d", isType, i)
		}
	}
}
