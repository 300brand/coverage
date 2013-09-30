package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
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
	logger.Trace.Printf("GetArticle: called %+v", query)
	switch v := query.(type) {
	case bson.ObjectId:
		query = bson.M{"_id": v}
	}
	err = m.C.Articles.Find(query).One(a)
	return
}

func (m *Mongo) UpdateArticle(a *coverage.Article) (err error) {
	logger.Trace.Printf("UpdateArticle: called %s", a.ID.Hex())
	a.Updated = time.Now()
	_, err = m.C.Articles.UpsertId(a.ID, a)
	if err != nil {
		logger.Error.Printf("UpdateArticle: %s", err)
		return
	}
	return
}
