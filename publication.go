package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Publication struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string
	URL         *url.URL
	NumFeeds    int64
	NumArticles int64
	Log         logger.Entries
	Added       time.Time
	Updated     time.Time
	LastCheck   time.Time
}

func NewPublication() (p *Publication) {
	p = &Publication{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	p.Log.Debug("Created: %s", p.ID.Hex())
	return
}
