package body

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

type Div string
type Divs []Div

func GetBody(b []byte) (body []byte, err error) {
	doc, err := gokogiri.ParseHtml(b)
	if err != nil {
		return
	}
	html := doc.Root().FirstChild()
	divs, err := GetPDivs(html)
	if err != nil {
		return
	}
	body = []byte(fmt.Sprintf("%d", len(divs)))
	return
}

// $divs = $xpath->query('//div[p]');
func GetPDivs(n xml.Node) (divs Divs, err error) {
	results, err := n.Search("//div[p]")
	if err != nil {
		return
	}
	fmt.Printf("%d\n", len(results))
	divs = make(Divs, len(results))
	return
}

// $divs = $xpath->query('//div[br/following-sibling::*[1][self::br]]');
//func GetBrBrDivs(doc *html.HtmlDocument) (divs Divs) {
//}
