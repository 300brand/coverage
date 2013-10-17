package normalizer

import (
	"errors"
	"fmt"
	"github.com/300brand/coverage/feed/parser/atom"
	"github.com/300brand/coverage/feed/parser/decoder"
	"github.com/300brand/coverage/feed/parser/rdf"
	"github.com/300brand/coverage/feed/parser/rss"
	"net/url"
	"strings"
	"time"
)

type Default struct {
	Title    string
	Articles []Article
}

type Article struct {
	Published time.Time
	Title     string
	URL       *url.URL
}

func (d *Default) Normalize(doc decoder.Decoder) (err error) {
	switch v := doc.(type) {
	case *atom.Doc:
		err = d.normalizeAtom(v)
	case *rss.Doc:
		err = d.normalizeRSS(v)
	case *rdf.Doc:
		err = d.normalizeRDF(v)
	default:
		errors.New("Unknown Decoder type")
	}
	return
}

func (d *Default) normalizeAtom(doc *atom.Doc) (err error) {
	d.Title = doc.Title
	for i, e := range doc.Entry {
		if len(e.Link) == 0 {
			errors.New(fmt.Sprintf("No links found for entry [%d] in %+v", i, e))
			continue
		}

		u, err := url.Parse(strings.TrimSpace(e.Link[0].Href))
		if err != nil {
			errors.New(fmt.Sprintf("Invalid URL [%s]: %v", u, err))
			continue
		}

		d.Articles = append(d.Articles, Article{
			Published: e.Updated,
			Title:     e.Title,
			URL:       u,
		})
	}
	return
}

func (d *Default) normalizeRDF(doc *rdf.Doc) (err error) {
	d.Title = doc.Channel.Title
	for i, item := range doc.Item {
		if item.Link == "" {
			errors.New(fmt.Sprintf("Empty link found for entry [%d] in %+v", i, item))
			continue
		}

		u, err := url.Parse(strings.TrimSpace(item.Link))
		if err != nil {
			errors.New(fmt.Sprintf("Invalid URL [%s]: %v", u, err))
			continue
		}

		d.Articles = append(d.Articles, Article{
			Published: item.Date.Time(),
			Title:     item.Title,
			URL:       u,
		})
	}
	return
}

func (d *Default) normalizeRSS(doc *rss.Doc) (err error) {
	d.Title = doc.Channel.Title
	for i, item := range doc.Channel.Item {
		if item.Link == "" {
			errors.New(fmt.Sprintf("Empty link found for entry [%d] in %+v", i, item))
			continue
		}

		u, err := url.Parse(strings.TrimSpace(item.Link))
		if err != nil {
			errors.New(fmt.Sprintf("Invalid URL [%s]: %v", u, err))
			continue
		}

		d.Articles = append(d.Articles, Article{
			Published: item.PubDate.Time(),
			Title:     item.Title,
			URL:       u,
		})
	}
	return
}
