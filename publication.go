package coverage

import (
	"net/url"
)

type Publication struct {
	Title    string
	Homepage url.URL
	TLD      string
	Feeds    []Feed
	Logs     LogEntries
}
