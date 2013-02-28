package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"strings"
	"testing"
)

func TestBlockElement(t *testing.T) {
	test := `
	<!DOCTYPE html>
	<html data-valid="false">
		<head data-valid="false"></head>
		<body data-valid="false">
			<a data-valid="false"></a>
			<article data-valid="true"></article>
			<div data-valid="true"></div>
			<footer data-valid="true"></footer>
			<header data-valid="true"></header>
			<input data-valid="false">
			<nav data-valid="true"></nav>
			<section data-valid="true"></section>
			<span data-valid="false"></span>
			<table data-valid="false">
				<tr data-valid="false">
					<td data-valid="true"></td>
				</tr>
			</table>
		</body>
	</html>
	`
	testElements(t, test, BlockElement)
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

// Fetches all child nodes within a node, flattening them into a single array
func getNodes(n *html.Node) (nodes []*html.Node) {
	for c := n; c != nil; c = c.NextSibling {
		nodes = append(nodes, c)
		if c.FirstChild != nil {
			nodes = append(nodes, getNodes(c.FirstChild)...)
		}
	}
	return
}

// Gets the value for an attribute on an element
func getAttribute(n *html.Node, q string) string {
	for _, a := range n.Attr {
		if a.Key == q {
			return a.Val
		}
	}
	return ""
}

// Takes a string of HTML where elements for testing have the attribute
// data-valid. The value for data-valid should be either "true" or "false" to
// match the outcome of the filter when applied to the tag.
func testElements(t *testing.T, s string, f Filter) {
	r := strings.NewReader(s)
	doc, err := html.Parse(r)
	if err != nil {
		t.Error(err)
	}
	nodes := getNodes(doc)
	for _, node := range nodes {
		val := getAttribute(node, "data-valid")
		// Only work on elements where the data-valid attribute exists
		if val == "" {
			continue
		}
		valid := val == "true"
		if f(node) != valid {
			t.Errorf("Expected %v for %s", valid, node.Data)
		}
	}
}
