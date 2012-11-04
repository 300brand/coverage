package atom

import (
	"encoding/xml"
	"git.300brand.com/coverage/logger"
	"git.300brand.com/coverage/parser"
	"net/url"
	"time"
)

// Types below from: <Go source>/src/pkg/encoding/xml/read_test.go
type Atom struct {
	feed    parser.Feed
	XMLName xml.Name  `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated,attr"`
	Author  Person    `xml:"author"`
	Entry   []Entry   `xml:"entry"`
}

type Entry struct {
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated"`
	Author  Person    `xml:"author"`
	Summary Text      `xml:"summary"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type Person struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
	//InnerXML string `xml:",innerxml"` // Not necessary
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

// Verify interface implementation at compile-time
var _ parser.Decoder = &Atom{}

func (atom *Atom) Decode(data []byte) (err error) {
	atom.feed = parser.Feed{}
	if err = xml.Unmarshal(data, atom); err != nil {
		return
	}
	atom.feed.Title = atom.Title
	for i, e := range atom.Entry {
		if len(e.Link) == 0 {
			logger.Warnf("No links found for entry [%d] in %+v", i, e)
			continue
		}

		url, err := url.Parse(e.Link[0].Href)
		if err != nil {
			logger.Warnf("Invalid URL [%s]: %v", url, err)
			continue
		}

		atom.feed.Articles = append(atom.feed.Articles, parser.Article{
			Published: e.Updated,
			Title:     e.Title,
			URL:       *url,
		})
	}
	return
}

func (atom Atom) Feed() parser.Feed {
	return atom.feed
}
