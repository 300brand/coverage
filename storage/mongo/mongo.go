package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	ArticleCollection = "Articles"
)

type Mongo struct {
	DBName    string
	URL       string
	db        *mgo.Database
	s         *mgo.Session
	cArticles *mgo.Collection
}

var indexes = map[string][]mgo.Index{
	ArticleCollection: {
		mgo.Index{
			Key:        []string{"url"},
			Background: true,
			DropDups:   true,
			Sparse:     false,
			Unique:     true,
		},
	},
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
	m.cArticles = m.db.C(ArticleCollection)
	err = m.EnsureIndexes()
	return
}

func (m *Mongo) EnsureIndexes() (err error) {
	for name, indexSet := range indexes {
		for _, index := range indexSet {
			if err = m.db.C(name).EnsureIndex(index); err != nil {
				return
			}
		}
	}
	return
}

func (m *Mongo) GetArticle(query interface{}) (a *coverage.Article, err error) {
	a = &coverage.Article{}
	err = m.cArticles.Find(query).One(a)
	return
}

func (m *Mongo) UpdateArticle(a *coverage.Article) (err error) {
	_, err = m.cArticles.UpsertId(a.ID, a)
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

func (m *Mongo) storeFile(prefix string, file *coverage.File) (err error) {
	var f *mgo.GridFile

	// .fs suffix helps the mongofiles access files
	gfs := m.db.GridFS(prefix + ".fs")

	// Save file
	if f, err = gfs.Create(file.Name); err != nil {
		return
	}
	f.SetContentType(file.ContentType)
	f.Write(file.Data)
	if err = f.Close(); err != nil {
		return
	}

	// Remove previous files
	query := gfs.Find(bson.M{
		"_id": bson.M{
			"$ne": f.Id(),
		},
		"filename": file.Name,
	})
	iter := query.Iter()
	for gfs.OpenNext(iter, &f) {
		gfs.RemoveId(f.Id())
	}
	return iter.Err()
}
