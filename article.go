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
	Published   time.Time
	ProperNames ProperNames
	Logs        LogEntries
}
