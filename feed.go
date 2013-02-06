package coverage

import (
	"net/url"
	"time"
)

type Feed struct {
	Title     string
	URL       url.URL
	LastCheck time.Time
	Articles  []Article
	Logs      LogEntries
}
