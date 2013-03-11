package coverage

import (
	. "git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	URL       *url.URL
	HTML      []byte `bson:"-"`
	Body      Body   `bson:"-"`
	Logs      LogEntries
	Added     time.Time
	Updated   time.Time
	LastCheck time.Time
	Published time.Time
}

func NewArticle() (a *Article) {
	a = &Article{
		ID: bson.NewObjectId(),
	}
	a.Added = time.Now()
	return
}

func (a *Article) Files() []File {
	return []File{
		{
			Name:        a.ID.Hex() + ".html",
			ContentType: "text/html",
			Data:        a.HTML,
		},
		{
			Name:        a.ID.Hex() + ".body.html",
			ContentType: "text/html",
			Data:        a.Body.HTML,
		},
		{
			Name:        a.ID.Hex() + ".body.text",
			ContentType: "text/plain",
			Data:        a.Body.Text,
		},
	}
}

func (a *Article) Modified() {
	a.Updated = time.Now()
}
