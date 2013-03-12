package coverage

import (
	"fmt"
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Publication struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	URL       *url.URL
	Feeds     []bson.ObjectId
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

func (p *Publication) AddFeed(f *Feed) error {
	for _, fId := range p.Feeds {
		if fId.Hex() == f.ID.Hex() {
			return p.Log.Error(fmt.Errorf("Duplicate feed: %s [%s]", f.ID.Hex(), f.URL))
		}
	}
	p.Feeds = append(p.Feeds, f.ID)
	p.Log.Debug("Added Feed: %s [%s]", f.ID.Hex(), f.URL)
	return nil
}
