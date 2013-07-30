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
	selector := bson.M{
		"_id": coverage.KeywordId{
			Date:    id.Date.Truncate(24 * time.Hour),
			Keyword: id.Keyword,
		},
	}
	return m.C.Keywords.Find(selector).One(kw)
}
