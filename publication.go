package coverage

import (
	"labix.org/v2/mgo/bson"
	"net/url"
)

type Publication struct {
	ID       bson.ObjectId
	Title    string
	Homepage url.URL
	TLD      string
	Feeds    []Feed
	FeedIDs  []bson.ObjectId
	Logs     LogEntries
}
