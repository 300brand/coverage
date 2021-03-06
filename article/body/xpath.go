package body

import (
	"fmt"
	"github.com/300brand/coverage"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"strings"
)

func XPath(in []byte, xpaths []string, body *coverage.Body) (err error) {
	// Stand up document tree
	doc, err := gokogiri.ParseHtml(in)
	if err != nil {
		return
	}
	defer doc.Free()

	// Document root
	root := doc.Root()
	defer root.Free()

	// Use remove-paths; store search-paths for later
	search := make([]string, 0, len(xpaths))
	for _, path := range xpaths {
		if path[0] != '-' {
			search = append(search, path)
			continue
		}
		if err = xpathRemove(root, path[1:]); err != nil {
			return
		}
	}

	// Search until something pops out
	var bodyNode xml.Node
	for _, path := range search {
		if err = xpathSearch(root, path, &bodyNode); err != nil {
			return
		}
		if bodyNode != nil {
			break
		}
	}
	if bodyNode == nil {
		return fmt.Errorf("No body node found")
	}
	// Clean out attributes
	xpathRemoveAttrs(bodyNode)
	body.HTML = []byte(bodyNode.String())
	// if false {
	// 	body.Text = []byte(filter.TranslateString(html.UnescapeString(bodyNode.Content())))
	// }
	body.Text = []byte(strings.Trim(bodyNode.Content(), " \t\r\n") + "\n")
	return
}

func xpathRemove(root xml.Node, path string) (err error) {
	// Find nodes
	nodes, err := root.Search(path)
	if err != nil {
		return
	}

	// Remove nodes
	for _, node := range nodes {
		node.Remove()
	}
	return
}

func xpathRemoveAttrs(root xml.Node) (err error) {
	for name, attr := range root.Attributes() {
		if name == "href" {
			continue
		}
		attr.Remove()
	}
	for c := root.FirstChild(); c != nil; c = c.NextSibling() {
		if c.NodeType() == xml.XML_TEXT_NODE {
			continue
		}
		xpathRemoveAttrs(c)
	}
	return
}

func xpathSearch(root xml.Node, path string, body *xml.Node) (err error) {
	// Do search
	nodes, err := root.Search(path)
	if err != nil {
		return
	}
	if len(nodes) == 0 {
		return fmt.Errorf("No nodes found")
	}
	// Temporarily reference just the first node
	node := nodes[0]
	// Set the outer block to just a div tag
	node.SetName("div")

	if err = BrBr2P(node); err != nil {
		return
	}

	*body = node
	return
}
