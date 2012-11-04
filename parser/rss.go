package parser

import (
	"encoding/xml"
	"time"
)

type RSS struct{}

type rss_Feed struct {
	XMLName xml.Name    `xml:"rss"`
	Channel rss_Channel `xml:"channel"`
}

type rss_Channel struct {
	Category       []string  `xml:"category,omitempty"`       // Optional. Defines one or more categories for the feed
	Cloud          string    `xml:"cloud,omitempty"`          // Optional. Register processes to be notified immediately of updates of the feed
	Copyright      string    `xml:"copyright,omitempty"`      // Optional. Notifies about copyrighted material
	Description    string    `xml:"description"`              // Required. Describes the channel
	Docs           string    `xml:"docs,omitempty"`           // Optional. Specifies an URL to the documentation of the format used in the feed
	Generator      string    `xml:"generator,omitempty"`      // Optional. Specifies the program used to generate the feed
	Image          string    `xml:"image,omitempty"`          // Optional. Allows an image to be displayed when aggregators present a feed
	Language       string    `xml:"language,omitempty"`       // Optional. Specifies the language the feed is written in
	LastBuildDate  time.Time `xml:"lastBuildDate,omitempty"`  // Optional. Defines the last-modified date of the content of the feed
	Link           string    `xml:"link"`                     // Required. Defines the hyperlink to the channel
	ManagingEditor string    `xml:"managingEditor,omitempty"` // Optional. Defines the e-mail address to the editor of the content of the feed
	PubDate        time.Time `xml:"pubDate,omitempty"`        // Optional. Defines the last publication date for the content of the feed
	Rating         string    `xml:"rating,omitempty"`         // Optional. The PICS rating of the feed
	SkipDays       string    `xml:"skipDays,omitempty"`       // Optional. Specifies the days where aggregators should skip updating the feed
	SkipHours      string    `xml:"skipHours,omitempty"`      // Optional. Specifies the hours where aggregators should skip updating the feed
	TextInput      string    `xml:"textInput,omitempty"`      // Optional. Specifies a text input field that should be displayed with the feed
	Title          string    `xml:"title"`                    // Required. Defines the title of the channel
	Ttl            string    `xml:"ttl,omitempty"`            // Optional. Specifies the number of minutes the feed can stay cached before refreshing it from the source
	WebMaster      string    `xml:"webMaster,omitempty"`      // Optional. Defines the e-mail address to the webmaster of the feed
}

func init() {
	decoders["RSS"] = RSS{}
}

func (a RSS) Decode(data []byte) (feed Feed, err error) {
	v := &rss_Feed{}
	if err = xml.Unmarshal(data, v); err != nil {
		return
	}
	feed.Title = v.Channel.Title
	/*
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
	*/
	return
}
