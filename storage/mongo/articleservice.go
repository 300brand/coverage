package mongo

import (
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
	"github.com/jbaikge/logger"
	"labix.org/v2/mgo/bson"
	"time"
)

type ArticleService struct {
	m *Mongo
}

const ArticleCollection = "Articles"

var _ service.ArticleService = &ArticleService{}

func NewArticleService(m *Mongo) *ArticleService {
	return &ArticleService{m: m}
}

func (s *ArticleService) Update(a *coverage.Article) error {
	a.Log.Service("mongo.ArticleService")
	return s.m.UpdateArticle(a)
}

func (m *Mongo) GetArticle(query interface{}, a *coverage.Article) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("GetArticle: called %+v", query)
	switch v := query.(type) {
	case bson.ObjectId:
		query = bson.M{"_id": v}
	}
	err = c.Articles.Find(query).One(a)
	return
}

func (m *Mongo) UpdateArticle(a *coverage.Article) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("UpdateArticle: called %s", a.ID.Hex())
	a.Updated = time.Now()
	_, err = c.Articles.UpsertId(a.ID, a)
	if err != nil {
		logger.Error.Printf("UpdateArticle: %s", err)
		return
	}
	return
}
