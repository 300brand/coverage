package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Publication struct {
	ID       bson.ObjectId
	Title    string
	Homepage url.URL
	TLD      string
	Feeds    []Feed
	FeedIDs  []bson.ObjectId
	Logs     logger.Entries
	Times    struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
	}
}
