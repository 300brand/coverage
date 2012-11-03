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
