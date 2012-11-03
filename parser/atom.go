package parser

import (
	"encoding/xml"
)

type Atom struct {
}

type atom_feed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string   `xml:"title"`
}

func init() {
	decoders["Atom"] = Atom{}
}

func (a Atom) Decode(data []byte) (feed Feed, err error) {
	v := &atom_feed{}
	if err = xml.Unmarshal(data, v); err != nil {
		return
	}
	feed.Title = v.Title
	return
}
