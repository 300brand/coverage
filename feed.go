package coverage

import (
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Feed struct {
	ID        bson.ObjectId
	Title     string
	URL       url.URL
	LastCheck time.Time
	Articles  []Article
	Logs      LogEntries
}
