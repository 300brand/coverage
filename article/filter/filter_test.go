package filter

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"strings"
	"testing"
)

func TestBlockElement(t *testing.T) {
	test := `<!DOCTYPE html>
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
	</html>`
	testElements(t, test, BlockElement)
}

func TestComment(t *testing.T) {
	// How many comments are in the test code below:
	expect := 1
	test := `<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<!-- COMMENT -->
		</body>
	</html>`
	for _, n := range getStringNodes(t, test) {
		if Comment(n) {
			expect--
		}
	}
	if expect > 0 {
		t.Errorf("Couldn't find %d comment(s)", expect)
	} else if expect < 0 {
		t.Errorf("Found %d too many comments", expect*-1)
	}
}

func TestEmpty(t *testing.T) {
	test := `<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<br id="1" data-valid="false">
			<div id="2" data-valid="true"></div>
			<div id="3" data-valid="false"><div id="4" data-valid="true"></div></div>
			<div id="5" data-valid="true">    </div>
			<div id="6" data-valid="true">
			</div>
			<div id="7" data-valid="false">
				<div id="8" data-valid="true"></div>
			</div>
			<div id="9" data-valid="false">Text</div>
		</body>
	</html>`
	testElements(t, test, Empty)
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

// Gets the value for an attribute on an element
func getAttribute(n *html.Node, q string) string {
	for _, a := range n.Attr {
		if a.Key == q {
			return a.Val
		}
	}
	return ""
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

func getStringNodes(t *testing.T, s string) []*html.Node {
	r := strings.NewReader(s)
	doc, err := html.Parse(r)
	if err != nil {
		t.Error(err)
	}
	return getNodes(doc)
}

// Takes a string of HTML where elements for testing have the attribute
// data-valid. The value for data-valid should be either "true" or "false" to
// match the outcome of the filter when applied to the tag.
func testElements(t *testing.T, s string, f Filter) {
	nodes := getStringNodes(t, s)
	for _, node := range nodes {
		val := getAttribute(node, "data-valid")
		// Only work on elements where the data-valid attribute exists
		if val == "" {
			continue
		}
		valid := val == "true"
		if f(node) != valid {
			t.Errorf("Expected %v for %s[%s]", valid, node.Data, getAttribute(node, "id"))
		}
	}
}
