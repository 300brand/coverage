package coverage

import (
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID          bson.ObjectId
	Title       string
	URL         url.URL
	ProperNames ProperNames
	Logs        LogEntries
	Times       struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
		Published time.Time
	}
}
