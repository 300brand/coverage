package coverage

import (
	"github.com/300brand/coverage/logger"
	"github.com/300brand/coverage/merger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID            bson.ObjectId `bson:"_id"`
	FeedId        bson.ObjectId `json:",omitempty"`
	PublicationId bson.ObjectId `json:",omitempty"`
	Title         string
	Author        string
	URL           *url.URL
	Text          Text
	Added         time.Time
	Updated       time.Time
	LastCheck     time.Time
	Published     time.Time
	Log           logger.Entries
	Changelog     merger.Changelog
}

func NewArticle() (a *Article) {
	a = &Article{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	a.Log.Debug("Created: %s", a.ID.Hex())
	return
}

func (a *Article) Modified(fields ...string) {
	a.Updated = time.Now()
	fields = append(fields, "Updated")
	a.Changelog.Changed(fields...)
}
