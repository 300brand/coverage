package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
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

func (m *Mongo) GetArticle(query interface{}) (a *coverage.Article, err error) {
	a = &coverage.Article{}
	err = m.C.Articles.Find(query).One(a)
	return
}

func (m *Mongo) UpdateArticle(a *coverage.Article) (err error) {
	a.Updated = time.Now()
	_, err = m.C.Articles.UpsertId(a.ID, a)
	if err != nil {
		return
	}
	return
}
