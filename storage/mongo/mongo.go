package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo"
)

const (
	ArticleCollection = "Articles"
)

type Mongo struct {
	DBName string
	URL    string
	db     *mgo.Database
	s      *mgo.Session
}

func New(url, dbName string) *Mongo {
	return &Mongo{
		URL:    url,
		DBName: dbName,
	}
}

func (m *Mongo) Close() {
	m.s.Close()
}

func (m *Mongo) Connect() (err error) {
	m.s, err = mgo.Dial(m.URL)
	if err != nil {
		return
	}
	m.db = m.s.DB(m.DBName)
	return
}

func (m *Mongo) GetArticle(query interface{}) (a *coverage.Article, err error) {
	a = &coverage.Article{}
	err = m.db.C(ArticleCollection).Find(query).One(a)
	return
}

func (m *Mongo) UpdateArticle(a *coverage.Article) (err error) {
	c := m.db.C(ArticleCollection)
	if _, err = c.UpsertId(a.ID, a); err != nil {
		return
	}
	for _, f := range a.Files() {
		if err = m.storeFile(ArticleCollection, f.Name, f.Data); err != nil {
			return
		}
	}
	return
}

func (m *Mongo) storeFile(prefix, name string, data []byte) (err error) {
	// .fs suffix helps the mongofiles access files
	g := m.db.GridFS(prefix + ".fs")
	f, err := g.Create(name)
	if err != nil {
		return
	}
	f.Write(data)
	return f.Close()
}
