package multipage

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
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

func Pages(start string, html []byte) (links []Link, err error) {
	if links, err = FindLinks(html); err != nil {
		return
	}

	if len(links) == 0 {
		return
	}

	base, err := url.Parse(start)
	if err != nil {
		return
	}

	for i := len(links) - 1; i >= 0; i-- {
		resolved := base.ResolveReference(links[i].Url)
		if resolved.String() == base.String() {
			links = append(links[:i], links[i+1:]...)
		}
		*links[i].Url = *resolved
	}
	return
}
