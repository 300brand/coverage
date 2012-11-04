package parser

import (
	"encoding/xml"
	"git.300brand.com/coverage/logger"
	"net/url"
	"time"
)

type Atom struct{}

// Types below from: <Go source>/src/pkg/encoding/xml/read_test.go
type atom_Feed struct {
	XMLName xml.Name     `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string       `xml:"title"`
	Id      string       `xml:"id"`
	Link    []atom_Link  `xml:"link"`
	Updated time.Time    `xml:"updated,attr"`
	Author  atom_Person  `xml:"author"`
	Entry   []atom_Entry `xml:"entry"`
}

type atom_Entry struct {
	Title   string      `xml:"title"`
	Id      string      `xml:"id"`
	Link    []atom_Link `xml:"link"`
	Updated time.Time   `xml:"updated"`
	Author  atom_Person `xml:"author"`
	Summary atom_Text   `xml:"summary"`
}

type atom_Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type atom_Person struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
	//InnerXML string `xml:",innerxml"` // Not necessary
}

type atom_Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

func init() {
	decoders["Atom"] = Atom{}
}

func (a Atom) Decode(data []byte) (feed Feed, err error) {
	v := &atom_Feed{}
	if err = xml.Unmarshal(data, v); err != nil {
		return
	}
	feed.Title = v.Title
	for i, e := range v.Entry {
		if len(e.Link) == 0 {
			logger.Warnf("No links found for entry [%d] in %+v", i, e)
			continue
		}

		url, err := url.Parse(e.Link[0].Href)
		if err != nil {
			logger.Warnf("Invalid URL [%s]: %v", url, err)
			continue
		}

		feed.Articles = append(feed.Articles, Article{
			Published: e.Updated,
			Title:     e.Title,
			URL:       *url,
		})
	}
	return
}
