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
	HTML  []byte `bson:"-"`
	Body  Body   `bson:"-"`
	Logs  logger.Entries
	Times struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
		Published time.Time
	}
}

type ArticleFile struct {
	Name string
	Data []byte
}

func NewArticle() (a *Article) {
	a = &Article{
		ID: bson.NewObjectId(),
	}
	a.Times.Added = time.Now()
	return
}

func (a *Article) Files() []ArticleFile {
	return []ArticleFile{
		{
			Name: a.ID.Hex() + ".html",
			Data: a.HTML,
		},
		{
			Name: a.ID.Hex() + ".body.html",
			Data: a.Body.HTML,
		},
		{
			Name: a.ID.Hex() + ".body.text",
			Data: a.Body.Text,
		},
	}
}

func (a *Article) Modified() {
	a.Times.Updated = time.Now()
}
