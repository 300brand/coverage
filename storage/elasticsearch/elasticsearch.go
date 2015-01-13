package elasticsearch

import (
	"github.com/300brand/coverage"
	"github.com/300brand/logger"
	elastic "github.com/mattbaird/elastigo/lib"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
)

type ElasticSearch struct {
	Host string
	Conn *elastic.Conn
}

type ElasticArticle struct {
	PublicationId bson.ObjectId
	FeedId        bson.ObjectId
	Title         string
	Author        string
	Body          string
	URL           string
	Published     time.Time
}

func New(host string) (es *ElasticSearch) {
	es = &ElasticSearch{Host: host}
	logger.Trace.Printf("Connecting to %s", host)
	es.Conn = elastic.NewConn()
	hosts := strings.Split(host, ";")
	es.Conn.SetHosts(hosts)
	return
}

func (es *ElasticSearch) SaveArticle(a *coverage.Article) (err error) {
	logger.Trace.Printf("SaveArticle %s", a.ID.Hex())
	ea := ElasticArticle{
		PublicationId: a.PublicationId,
		FeedId:        a.FeedId,
		Title:         a.Title,
		Author:        a.Author,
		Body:          string(a.Text.Body.Text),
		URL:           a.URL,
		Published:     a.Published,
	}
	args := map[string]interface{}{
		"timestamp": a.Added,
	}
	_, err = es.Conn.Index("articles", "article", a.ID.Hex(), args, ea)
	if err != nil {
		logger.Error.Printf("SaveArticle: %s", err)
	}
	return
}
