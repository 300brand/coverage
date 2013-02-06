package coverage

import (
	"net/url"
	"time"
)

type Article struct {
	Title       string
	URL         url.URL
	Published   time.Time
	ProperNames ProperNames
	Logs        LogEntries
}
