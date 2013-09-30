package mongo

import (
	"git.300brand.com/coverage"
	"github.com/jbaikge/logger"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

const KeywordCollection = "Keywords"

func (m *Mongo) AddKeywords(a *coverage.Article) (err error) {
	logger.Trace.Printf("AddKeywords: called %s", a.ID.Hex())
	var word string
	selector := bson.M{
		"_id": bson.M{
			"date":    dtoi(a.Published),
			"keyword": &word, // Using bson.M's instead of KeywordId because of this
		},
	}
	change := bson.M{"$addToSet": bson.M{"articles": a.ID}}
	for _, word = range a.Text.Words.Keywords {
		if _, err = m.C.Keywords.Upsert(selector, change); err != nil {
			return
		}
	}
	return
}

func (m *Mongo) GetKeyword(id *coverage.KeywordId, kw *coverage.Keyword) (err error) {
	logger.Trace.Printf("GetKeyword: called %+v", id)
	log.Printf("Mongo.GetKeyword id: %+v", id)
	selector := bson.M{
		"_id": bson.M{
			"date":    dtoi(id.Date),
			"keyword": id.Keyword,
		},
	}
	log.Printf("Mongo.GetKeyword: %+v", selector)
	return m.C.Keywords.Find(selector).One(kw)
}

func dtoi(t time.Time) int {
	y, m, d := t.Date()
	return y*1e4 + int(m)*1e2 + d
}
