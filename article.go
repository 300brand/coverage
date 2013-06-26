package coverage

import (
	"git.300brand.com/coverage/logger"
	"git.300brand.com/coverage/merger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID        bson.ObjectId `bson:"_id"`
	FeedId    bson.ObjectId `bson:",omitempty"`
	Title     string
	URL       *url.URL
	Text      Text
	Added     time.Time
	Updated   time.Time
	LastCheck time.Time
	Published time.Time
	Log       logger.Entries
	Changelog merger.Changelog
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
