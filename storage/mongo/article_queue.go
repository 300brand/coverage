package mongo

import (
	"github.com/300brand/coverage"
	"github.com/300brand/logger"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const ArticleQueueCollection = "ArticleQ"

func (m *Mongo) ArticleQueueAdd(a *coverage.Article) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("ArticleQueueAdd: called %s", a.ID.Hex())
	return c.ArticleQ.Insert(a)
}

func (m *Mongo) ArticleQueueNext(a *coverage.Article) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("ArticleQueueNext: called")
	query := bson.M{
		"queue":   time.Now().Second() / (60 / 4),
		"dequeue": bson.M{"$lte": time.Now()},
		"tries":   bson.M{"$lt": 10},
	}
	change := mgo.Change{
		Remove:    false,
		ReturnNew: false,
		Update: bson.M{
			"$set": bson.M{
				"dequeue": time.Now().Add(time.Hour),
			},
			"$inc": bson.M{
				"tries": 1,
			},
		},
		Upsert: false,
	}
	logger.Trace.Printf("Query: %+v", query)
	logger.Trace.Printf("Change: %+v", change)
	_, err = c.ArticleQ.Find(query).Limit(1).Apply(change, a)
	if err != nil {
		return
	}
	logger.Trace.Printf("ArticleQueueNext: %s", a.ID.Hex())
	return
}

func (m *Mongo) ArticleQueueRemove(id bson.ObjectId) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("ArticleQueueRemove: called %s", id.Hex())
	return c.ArticleQ.RemoveId(id)
}
