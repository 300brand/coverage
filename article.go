package coverage

import (
	"git.300brand.com/coverage/article/body"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID    bson.ObjectId
	Title string
	URL   *url.URL
	HTML  []byte
	Body  body.Body
	Logs  logger.Entries
	Times struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
		Published time.Time
	}
}

func NewArticle(r *downloader.Response) (a *Article, err error) {
	a = &Article{
		ID: bson.NewObjectId(),
	}
	a.Times.Added = time.Now()
	// Temporarily commented out - friggin huge
	// a.HTML = r.Body

	if a.URL, err = url.Parse(r.RealURL); err != nil {
		return
	}

	if a.Body, err = body.GetBody(r.Body); err != nil {
		return
	}

	a.Times.Updated = time.Now()
	a.Times.LastCheck = time.Now()
	return
}
