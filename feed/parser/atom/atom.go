package atom

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"github.com/300brand/coverage/feed/parser/decoder"
	"time"
)

// Types below from: <Go source>/src/pkg/encoding/xml/read_test.go
type Doc struct {
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
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Person struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
	//InnerXML string `xml:",innerxml"` // Not necessary
}

type Text struct {
	Type string `xml:"type,attr"`
	Body string `xml:",chardata"`
}

func init() {
	decoder.RegisterDecoder("atom", &Doc{})
}

func (doc Doc) New() decoder.Decoder {
	return &Doc{}
}

func (doc *Doc) Decode(data []byte) error {
	d := xml.NewDecoder(bytes.NewReader(data))
	d.CharsetReader = charset.NewReader
	return d.Decode(doc)
}
