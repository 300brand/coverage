package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Feed struct {
	ID        bson.ObjectId
	Title     string
	URL       url.URL
	Articles  []Article `bson:-`
	Logs      logger.Entries
	Added     time.Time
	Updated   time.Time
	LastCheck time.Time
}
