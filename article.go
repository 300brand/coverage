package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID    bson.ObjectId `bson:"_id"`
	Title string
	URL   *url.URL
	HTML  []byte
	Body  Body
	Logs  logger.Entries
	Times struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
		Published time.Time
	}
}

func (a *Article) Modified() {
	a.Times.Updated = time.Now()
}

func NewArticle() (a Article) {
	a.ID = bson.NewObjectId()
	a.Times.Added = time.Now()
	return
}
