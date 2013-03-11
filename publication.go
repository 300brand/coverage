package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Publication struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	URL       *url.URL
	Feeds     []Feed
	Log       logger.Entries
	Added     time.Time
	Updated   time.Time
	LastCheck time.Time
}

func NewPublication() (p *Publication) {
	p = &Publication{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	p.Log.Debug("Created: %s", p.ID.Hex())
	return
}

func (p *Publication) AddFeed(f *Feed) {
	p.Feeds = append(p.Feeds, *f)
	p.Log.Debug("Added Feed: %s", f.ID.Hex())
}
