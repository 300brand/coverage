package coverage

import (
	"github.com/300brand/coverage/logger"
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
	NumReaders  int64
	Deleted     bool
	AuthorPaths []string
	BodyPaths   []string
	Log         logger.Entries
	Added       time.Time
	Updated     time.Time
}

func NewPublication() (p *Publication) {
	p = &Publication{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	return
}
