package body

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

type Block string
type Blocks []Block

func GetBody(b []byte) (body []byte, err error) {
	doc, err := gokogiri.ParseHtml(b)
	defer doc.Free()
	if err != nil {
		return
	}
	html := doc.Root()
	fmt.Println(html.Name())
	blocks, err := GetPBlocks(html)
	if err != nil {
		return
	}
	body = []byte(fmt.Sprintf("GetBody %q", blocks))
	return
}

func GetPBlocks(n xml.Node) (b Blocks, err error) {
	results, err := n.Search("//div[p]")
	if err != nil {
		return
	}
	for _, r := range results {
		b = append(b, Block(r.Content()))
	}
	return
}

// $divs = $xpath->query('//div[br/following-sibling::*[1][self::br]]');
//func GetBrBrDivs(doc *html.HtmlDocument) (divs Divs) {
//}
