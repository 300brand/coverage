package filter

import (
	"code.google.com/p/go.net/html"
	"strings"
	"testing"
)

func TestAside(t *testing.T) {
	test := `<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<div class="aside" valid="false"></div>
			<div valid="false">Text</div>
			<aside valid="true">
				<div valid="false"></div>
			</aside>
			<div valid="false">
				<aside valid="true"></aside>
			</div>
		</body>
	</html>`
	testElements(t, test, Aside)
}

func TestBlockElement(t *testing.T) {
	test := `<!DOCTYPE html>
	<html valid="false">
		<head valid="false"></head>
		<body valid="false">
			<a valid="false"></a>
			<article valid="true"></article>
			<div valid="true"></div>
			<footer valid="true"></footer>
			<header valid="true"></header>
			<input valid="false">
			<nav valid="true"></nav>
			<section valid="true"></section>
			<span valid="false"></span>
			<table valid="false">
				<tr valid="false">
					<td valid="true"></td>
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
			<br id="1" valid="false">
			<div id="2" valid="true"></div>
			<div id="3" valid="false"><div id="4" valid="true"></div></div>
			<div id="5" valid="true">    </div>
			<div id="6" valid="true">
			</div>
			<div id="7" valid="false">
				<div id="8" valid="true"></div>
			</div>
			<div id="9" valid="false">Text</div>
		</body>
	</html>`
	testElements(t, test, Empty)
}

func TestForm(t *testing.T) {
	test := `<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<form valid="true" action="login" method="post">
				<input type="text">
				<input type="password">
			</form>
		</body>
	</html>`
	testElements(t, test, Form)
}

func TestHead(t *testing.T) {
	test := `<!DOCTYPE html><html><head></head><body></body></html>`
	testElements(t, test, Head)
}

func TestScript(t *testing.T) {
	test := `<!DOCTYPE html>
	<html valid="false">
		<head valid="false">
			<script valid="true"></script>
			<script valid="true">
				document.write('mooo')
			</script>
		</head>
		<body valid="false">
			<script valid="true" src="./thing.js"></script>
		</body>
	</html>`
	testElements(t, test, Script)
}

func TestSpaces(t *testing.T) {
	expect := 7
	test := `<!DOCTYPE html><html><head></head><body>
		<div></div>
		<div>
			<h1>Test</h1>
			<h2> </h2>
		</div>
	</body></html>`
	for _, n := range getStringNodes(t, test) {
		if Spaces(n) {
			expect--
		}
	}
	if expect > 0 {
		t.Errorf("Couldn't find %d space node(s)", expect)
	} else if expect < 0 {
		t.Errorf("Found %d too many space nodes", expect*-1)
	}
}

func TestStyle(t *testing.T) {
	test := `<!DOCTYPE html>
	<html valid="false">
		<head valid="false">
			<style valid="true" type="text/css">
				body { color:#F00; }
			</style>
		</head>
		<body valid="false">
			<style valid="true"></style>
		</body>
	</html>`
	testElements(t, test, Style)
}

func TestText(t *testing.T) {
	expect := 8
	test := `<!DOCTYPE html><html><head></head><body>
		<div></div>
		<div>
			<h1>Test</h1>
			<h2> </h2>
		</div>
	</body></html>`
	for _, n := range getStringNodes(t, test) {
		if Text(n) {
			expect--
		}
	}
	if expect > 0 {
		t.Errorf("Couldn't find %d text node(s)", expect)
	} else if expect < 0 {
		t.Errorf("Found %d too many text nodes", expect*-1)
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
// valid. The value for valid should be either "true" or "false" to
// match the outcome of the filter when applied to the tag.
func testElements(t *testing.T, s string, f Filter) {
	nodes := getStringNodes(t, s)
	for _, node := range nodes {
		val := getAttribute(node, "valid")
		// Only work on elements where the valid attribute exists
		if val == "" {
			continue
		}
		valid := val == "true"
		if f(node) != valid {
			t.Errorf("Expected %v for %s[%s]", valid, node.Data, getAttribute(node, "id"))
		}
	}
}
