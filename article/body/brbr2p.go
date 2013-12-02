package body

import (
	"github.com/moovweb/gokogiri/xml"
	"strings"
)

func BrBr2P(root xml.Node) (err error) {
	// Drop double-br's to singles
	nodes, err := root.Search("//br")
	if err != nil {
		return
	}
	for _, node := range nodes {
		sibling := node.NextSibling()
		if sibling != nil && sibling.Name() == "br" {
			node.Remove()
		}
	}
	// Convert content before br's to p's
	var next, p xml.Node
	for node := root.FirstChild(); node != nil; node = next {
		next = node.NextSibling()

		if p == nil {
			p = newP(root)
			node.InsertBefore(p)
		}

		switch node.NodeType() {
		case xml.XML_TEXT_NODE:
			node.SetContent(" " + strings.TrimSpace(node.Content()) + " ")
			if node.Content() == "  " {
				node.Remove()
			} else {
				p.AddChild(node)
			}
		case xml.XML_COMMENT_NODE:
			node.Remove()
		case xml.XML_ELEMENT_NODE:
			switch node.Name() {
			case "p":
				continue
			case "br":
				node.SetName("p")
				cleanP(p)
				p = node
				continue
			}
			p.AddChild(node)
		}
	}
	cleanP(p)
	return
}

func cleanP(p xml.Node) {
	// Drop if there's nothing inside
	if p.CountChildren() == 0 {
		p.Remove()
		return
	}
	// Trim first text node's leading space
	if c := p.FirstChild(); c.NodeType() == xml.XML_TEXT_NODE {
		c.SetContent(strings.TrimLeft(c.Content(), " "))
	}
	// Trim last text node's trailing space
	if c := p.LastChild(); c.NodeType() == xml.XML_TEXT_NODE {
		c.SetContent(strings.TrimRight(c.Content(), " "))
	}
	p.InsertAfter("\n")
}

func newP(root xml.Node) (node xml.Node) {
	fragment, err := root.MyDocument().ParseFragment([]byte(`<p/>`), nil, xml.DefaultParseOption)
	if err != nil {
		panic(err)
	}
	return fragment.FirstChild()
}
