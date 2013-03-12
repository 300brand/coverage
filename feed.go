package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Feed struct {
	ID        bson.ObjectId `bson:"_id"`
	URL       *url.URL
	Content   []byte `bson:"-"`
	Log       logger.Entries
	Added     time.Time
	Updated   time.Time
	LastCheck time.Time
}

func NewFeed() (f *Feed) {
	f = &Feed{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	f.Log.Debug("Created: %s", f.ID.Hex())
	return
}
