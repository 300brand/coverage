package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"time"
)

const KeywordCollection = "Keywords"

func (m *Mongo) AddKeywords(a *coverage.Article) (err error) {
	var word string
	selector := bson.M{
		"_id": bson.M{
			"date":    a.Published.Truncate(24 * time.Hour),
			"keyword": &word,
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

func (m *Mongo) GetKeyword(kw *coverage.Keyword) (err error) {
	return m.C.Keywords.Find(bson.M{"_id": kw.Id}).One(kw)
}
