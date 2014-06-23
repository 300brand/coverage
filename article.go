package coverage

import (
	"github.com/300brand/coverage/logger"
	"github.com/300brand/coverage/merger"
	"labix.org/v2/mgo/bson"
	"time"
)

type Article struct {
	ID            bson.ObjectId `bson:"_id"`
	FeedId        bson.ObjectId `json:",omitempty"`
	PublicationId bson.ObjectId `json:",omitempty"`
	Title         string
	Author        string
	URL           string
	Text          Text
	Added         time.Time
	Updated       time.Time
	LastCheck     time.Time
	Published     time.Time
	PubDate       ArticlePubDate `bson:,omitempty`
	Queue         int            // Used for Article-queuing
	Tries         int            // Used for article-queuing
	Dequeue       time.Time      // Used for article-queuing
	Log           logger.Entries
	Changelog     merger.Changelog
}

type ArticlePubDate struct {
	Date int
}

var _ bson.Getter = new(Article)

func NewArticle() (a *Article) {
	a = &Article{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	// Queue: Seconds divided by 60-seconds divided by number of Mongo shards
	// (hard-coded because why not)
	a.Queue = a.Added.Second() / (60 / 4)
	a.Log.Debug("Created: %s", a.ID.Hex())
	return
}

func (a *Article) GetBSON() (out interface{}, err error) {
	y, m, d := a.Published.Date()
	a.PubDate.Date = y*1e4 + int(m)*1e2 + d
	out = interface{}(a)
	return
}

func (a *Article) Modified(fields ...string) {
	a.Updated = time.Now()
	fields = append(fields, "Updated")
	a.Changelog.Changed(fields...)
}
