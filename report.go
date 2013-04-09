package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Report struct {
	ID              bson.ObjectId `bson:"_id"`
	QueueId         uint64
	ObjectId        uint64
	PreviousResults []*url.URL
	Feeds           []*url.URL `bson:"-"`
	FeedIds         []bson.ObjectId
	Log             logger.Entries
	Added           time.Time
	Updated         time.Time
	Summaries       []Summary
	DateBounds      struct {
		Start time.Time
		End   time.Time
	}
	Phrases struct {
		Include []string
		Exclude []string
	}
}

func NewReport() (r *Report) {
	return &Report{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
}
