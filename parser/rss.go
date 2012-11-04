package parser

import (
	"encoding/xml"
)

type RSS struct{}

type rss_Feed struct {
	XMLName xml.Name    `xml:"rss"`
	Channel rss_Channel `xml:"channel"`
}

type rss_Channel struct {
	Title string `xml:"title"`
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
