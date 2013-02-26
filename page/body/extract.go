package body

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

type Body struct {
	HTML string
	Text string
}

type blockFunc func(xml.Node) ([]xml.Node, error)

var blockFuncs = []blockFunc{
	PBlocks,
	//BrBrBlocks,
}

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
	for i, block := range blocks {
		fmt.Printf("%d: %s\n\n", i, block)
	}
	return
}

func PBlocks(n xml.Node) ([]xml.Node, error) {
	return n.Search("//div[p]")
}

func BrBrBlocks(n xml.Node) ([]xml.Node, error) {
	return n.Search("//div[br/following-sibling::*[1][self::br]]")
}
