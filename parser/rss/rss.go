package rss

import (
	"encoding/xml"
	"git.300brand.com/coverage/logger"
	"git.300brand.com/coverage/parser"
	"git.300brand.com/coverage/strtotime"
	"net/url"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Category       []string `xml:"category,omitempty"`       // Optional. Defines one or more categories for the feed
	Cloud          string   `xml:"cloud,omitempty"`          // Optional. Register processes to be notified immediately of updates of the feed
	Copyright      string   `xml:"copyright,omitempty"`      // Optional. Notifies about copyrighted material
	Description    string   `xml:"description"`              // Required. Describes the channel
	Docs           string   `xml:"docs,omitempty"`           // Optional. Specifies an URL to the documentation of the format used in the feed
	Generator      string   `xml:"generator,omitempty"`      // Optional. Specifies the program used to generate the feed
	Image          string   `xml:"image,omitempty"`          // Optional. Allows an image to be displayed when aggregators present a feed
	Item           []Item   `xml:"item,omitempty"`           // Optional. Stories within the feed
	Language       string   `xml:"language,omitempty"`       // Optional. Specifies the language the feed is written in
	LastBuildDate  Time     `xml:"lastBuildDate,omitempty"`  // Optional. Defines the last-modified date of the content of the feed
	Link           string   `xml:"link"`                     // Required. Defines the hyperlink to the channel
	ManagingEditor string   `xml:"managingEditor,omitempty"` // Optional. Defines the e-mail address to the editor of the content of the feed
	PubDate        Time     `xml:"pubDate,omitempty"`        // Optional. Defines the last publication date for the content of the feed
	Rating         string   `xml:"rating,omitempty"`         // Optional. The PICS rating of the feed
	SkipDays       string   `xml:"skipDays,omitempty"`       // Optional. Specifies the days where aggregators should skip updating the feed
	SkipHours      string   `xml:"skipHours,omitempty"`      // Optional. Specifies the hours where aggregators should skip updating the feed
	TextInput      string   `xml:"textInput,omitempty"`      // Optional. Specifies a text input field that should be displayed with the feed
	Title          string   `xml:"title"`                    // Required. Defines the title of the channel
	Ttl            string   `xml:"ttl,omitempty"`            // Optional. Specifies the number of minutes the feed can stay cached before refreshing it from the source
	WebMaster      string   `xml:"webMaster,omitempty"`      // Optional. Defines the e-mail address to the webmaster of the feed
}

type Item struct {
	Author      string `xml:"author,omitempty"`    // Optional. Specifies the e-mail address to the author of the item
	Category    string `xml:"category,omitempty"`  // Optional. Defines one or more categories the item belongs to
	Comments    string `xml:"comments,omitempty"`  // Optional. Allows an item to link to comments about that item
	Description string `xml:"description"`         // Required. Describes the item
	Enclosure   string `xml:"enclosure,omitempty"` // Optional. Allows a media file to be included with the item
	Guid        string `xml:"guid,omitempty"`      // Optional. Defines a unique identifier for the item
	Link        string `xml:"link"`                // Required. Defines the hyperlink to the item
	PubDate     Time   `xml:"pubDate,omitempty"`   // Optional. Defines the last-publication date for the item
	Source      string `xml:"source,omitempty"`    // Optional. Specifies a third-party source for the item
	Title       string `xml:"title"`               // Required. Defines the title of the item
}

type Time string

func (t Time) Time() (ts time.Time) {
	ts, err := strtotime.Parse(string(t))
	if err != nil {
		logger.Warn(err)
	}
	return
}

// Verify interface implementation at compile-time
var _ parser.Decoder = &RSS{}

func (doc *RSS) Decode(data []byte) error {
	return xml.Unmarshal(data, doc)
}

func (doc RSS) Feed() (f parser.Feed) {
	f.Title = doc.Channel.Title
	for i, item := range doc.Channel.Item {
		if item.Link == "" {
			logger.Warnf("Empty link found for entry [%d] in %+v", i, item)
			continue
		}

		url, err := url.Parse(item.Link)
		if err != nil {
			logger.Warnf("Invalid URL [%s]: %v", url, err)
			continue
		}

		f.Articles = append(f.Articles, parser.Article{
			Published: item.PubDate.Time(),
			Title:     item.Title,
			URL:       *url,
		})
	}
	return
}
