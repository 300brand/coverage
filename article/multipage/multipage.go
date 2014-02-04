package multipage

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
	"log"
	"net/url"
	"strconv"
)

type Link struct {
	Num int
	Url *url.URL
}

var LinkXPath *xpath.Expression

func init() {
	path := `//a[number(.) = number(.) and not(contains(@class, "count"))]`
	if LinkXPath = xpath.Compile(path); LinkXPath == nil {
		panic("Invalid XPath: " + path)
	}
}

func FindLinks(html []byte) (links []Link, err error) {
	doc, err := gokogiri.ParseHtml(html)
	if err != nil {
		return
	}
	defer doc.Free()

	root := doc.Root()
	defer root.Free()

	nodes, err := root.Search(LinkXPath)
	if err != nil {
		return
	}
	log.Printf("Found %d <a> tags", len(nodes))
	links = make([]Link, 0, len(nodes))
	for i := range nodes {
		link := Link{}
		if link.Num, err = strconv.Atoi(nodes[i].InnerHtml()); err != nil {
			continue
		}
		href := nodes[i].Attr("href")
		if link.Url, err = url.Parse(href); err != nil {
			continue
		}
		links = append(links, link)
	}

	return
}
