package normalizer

import (
	"git.300brand.com/coverage/parser"
	"git.300brand.com/coverage/parser/atom"
	"git.300brand.com/coverage/parser/rdf"
	"git.300brand.com/coverage/parser/rss"
	"net/url"
	"time"
)

type Default struct {
	Title    string
	Articles []Article
}

type Article struct {
	Published time.Time
	Title     string
	URL       url.URL
}

func (d *Default) Normalize(doc Decoder) (err error) {
	switch v := doc.(type) {
	case atom.Doc:
		err = d.normalizeAtom(v)
	case rss.Doc:
		err = d.normalizeRSS(v)
	case rdf.Doc:
		err = d.normalizeRDF(v)
	default:
		errors.New("Unknown Decoder type")
	}
	return
}

func (d *Default) normalizeAtom(doc atom.Doc) (err error) {
	d.Title = doc.Title
	return
}

func (d *Default) normalizeRDF(doc rdf.Doc) (err error) {
	return
}

func (d *Default) normalizeRSS(doc rss.Doc) (err error) {
	d.Title = doc.Channel.Title
	return
}
