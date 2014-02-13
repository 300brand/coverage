package coverage

import (
	"github.com/300brand/coverage/logger"
	"labix.org/v2/mgo/bson"
	"time"
)

type Publication struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string
	URL         string
	NumFeeds    int64
	NumArticles int64
	NumReaders  int64
	Deleted     bool
	Log         logger.Entries
	Added       time.Time
	Updated     time.Time
	XPaths      struct{ Author, Body, Date []string }
}

func NewPublication() (p *Publication) {
	p = &Publication{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	return
}
