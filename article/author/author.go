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
		// Try again
		if author == "" {
			continue
		}
		// Success!
		return
	}
	return "", fmt.Errorf("No author found with %+v", xpaths)
}

func searchXpath(node xml.Node, path string) (match string, err error) {
	// Ensure normalize-space is used
	if !strings.Contains(path, "normalize-space") {
		path = fmt.Sprintf("normalize-space(%s)", path)
	}

	// Compile expression
	expression := xpath.Compile(path)
	if expression == nil {
		err = fmt.Errorf("Invalid XPath expression: %s", path)
		return
	}
	defer expression.Free()

	x := xpath.NewXPath(node.NodePtr())
	defer x.Free()

	if err = x.Evaluate(node.NodePtr(), expression); err != nil {
		return
	}

	match, err = x.ResultAsString()
	match = strings.TrimSpace(match)
	return
}
