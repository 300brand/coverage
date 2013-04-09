package coverage

import (
	"net/url"
	"time"
)

type Summary struct {
	PageId      uint64
	FeedId      uint64
	Author      string
	Body        string
	Url         *url.URL
	Title       string
	Publication string
	Published   time.Time
	New         bool
	Comments    uint64
	Facebook    uint64
	Linkedin    uint64
	Twitter     uint64
	Score       uint64
}
