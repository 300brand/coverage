package author

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"strings"
)

func Search(html []byte, xpaths []string) (author string, err error) {
	errs := make([]string, 0, len(xpaths))

	doc, err := gokogiri.ParseHtml(html)
	if err != nil {
		return
	}
	defer doc.Free()

	root := doc.Root()
	defer root.Free()

	for i, path := range xpaths {
		// Attempt
		author, err = searchXpath(root, path)
		if err != nil {
			errs = append(errs, fmt.Sprintf("[%d] %s", i, err.Error()))
		}
		// Success!
		return
	}
	return "", fmt.Errorf("%s", strings.Join(errs, "; "))
}

func searchXpath(node xml.Node, path string) (author string, err error) {
	// Compile expression
	expression := xpath.Compile(path)
	if expression == nil {
		err = fmt.Errorf("Invalid XPath expression: %s", path)
		return
	}
	defer expression.Free()
	// Search with expression on HTML root
	nodes, err := node.Search(expression)
	if err != nil {
		return
	}
	if len(nodes) == 0 {
		err = fmt.Errorf("No matches found")
	}
	if len(nodes) > 1 {
		// Trace more than one match, only using the first
	}
	author = nodes[0].Content()
	return
}
