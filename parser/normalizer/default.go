package parser

import (
	"net/url"
	"time"
)

type Feed struct {
	Title    string
	Articles []Article
}

type Article struct {
	Published time.Time
	Title     string
	URL       url.URL
}

func (f *Feed) Normalize(doc Decoder) (err error) {
	switch v := doc.(type) {
	case atom.Doc:
		err = f.normalizeAtom(v)
	case rss.Doc:
		err = f.normalizeRSS(v)
	case rdf.Doc:
		err = f.normalizeRDF(v)
	default:
		errors.New("Unknown Decoder type")
	}
	return
}

func (f *Feed) normalizeAtom(doc atom.Doc) (err error) {
	f.Title = doc.Title
	return
}

func (f *Feed) normalizeRDF(doc rdf.Doc) (err error) {
	return
}

func (f *Feed) normalizeRSS(doc rss.Doc) (err error) {
	f.Title = doc.Channel.Title
	return
}
