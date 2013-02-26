package body

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
)

type Body struct {
	HTML string
	Text string
}

type blockFunc func(xml.Node) ([]xml.Node, error)

var (
	blockFuncs = []blockFunc{
		PBlocks,
		BrBrBlocks,
	}
	xpathP    = xpath.Compile("//div[count(p) > 1]")
	xpathBrBr = xpath.Compile("//div[br/following-sibling::*[1][self::br]]")
)

func GetBody(in []byte) (b Body, err error) {
	doc, err := gokogiri.ParseHtml(in)
	defer doc.Free()
	if err != nil {
		return
	}
	html := doc.Root()
	blocks := []xml.Node{}
	for _, f := range blockFuncs {
		nodes, err := f(html)
		if err != nil {
			return b, err
		}
		blocks = append(blocks, nodes...)
	}
	for _, block := range blocks {
		content := block.Content()
		if len(content) > len(b.Text) {
			b.HTML = block.InnerHtml()
			b.Text = content
		}
	}
	return
}

func PBlocks(n xml.Node) ([]xml.Node, error) {
	return n.Search(xpathP)
}

func BrBrBlocks(n xml.Node) ([]xml.Node, error) {
	return n.Search(xpathBrBr)
}
