package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"labix.org/v2/mgo"
)

type ArticleService struct {
	m *Mongo
}

const ArticleCollection = "Articles"

var _ service.ArticleService = &ArticleService{}

func init() {
	indexes[ArticleCollection] = []mgo.Index{
		mgo.Index{
			Key:        []string{"url"},
			Background: true,
			DropDups:   true,
			Sparse:     false,
			Unique:     true,
		},
	}
}

func NewArticleService(m *Mongo) *ArticleService {
	return &ArticleService{m: m}
}

func (s *ArticleService) Update(a *coverage.Article) error {
	a.Log.Service("mongo.ArticleService")
	return s.m.UpdateArticle(a)
}

func (m *Mongo) GetArticle(query interface{}) (a *coverage.Article, err error) {
	a = &coverage.Article{}
	err = m.db.C(ArticleCollection).Find(query).One(a)
	return
}

func (m *Mongo) UpdateArticle(a *coverage.Article) (err error) {
	a.Log.Debug("mongo.UpdateArticle")
	_, err = m.db.C(ArticleCollection).UpsertId(a.ID, a)
	if err != nil {
		return
	}
	for _, f := range a.Files() {
		if err = m.storeFile(ArticleCollection, &f); err != nil {
			return
		}
	}
	return
}
