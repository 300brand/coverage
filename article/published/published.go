package published

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"strings"
	"time"
)

func Search(html []byte, xpaths []string) (date time.Time, err error) {
	errs := make([]string, 0, len(xpaths))

	doc, err := gokogiri.ParseHtml(html)
	if err != nil {
		return
	}
	defer doc.Free()

	root := doc.Root()
	defer root.Free()

	var match string
	for i, path := range xpaths {
		bits := strings.SplitN(path, "~~", 2)
		if len(bits) != 2 {
			errs = append(errs, fmt.Sprintf("[%d] Missing format: %s", i, path))
			continue
		}
		search, layout := bits[0], bits[1]
		// Attempt
		match, err = searchXpath(root, search)
		if err != nil {
			errs = append(errs, fmt.Sprintf("[%d] %s", i, err))
			continue
		}
		date, err = time.Parse(layout, match)
		if err != nil {
			errs = append(errs, fmt.Sprintf("[%d] %s", i, err))
			continue
		}
		// Try again
		if date.IsZero() {
			continue
		}
		// Success!
		return date, nil
	}
	if len(errs) > 0 {
		err = fmt.Errorf("Errors: %s", strings.Join(errs, "; "))
	} else {
		err = fmt.Errorf("No date found with %+v", xpaths)
	}
	return
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
	return
}
