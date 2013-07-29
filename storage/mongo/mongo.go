package mongo

import (
	"labix.org/v2/mgo"
)

type Mongo struct {
	C       collections
	Host    string
	Prefix  string
	Session *mgo.Session
}

type collections struct {
	Articles     *mgo.Collection
	Feeds        *mgo.Collection
	Keywords     *mgo.Collection
	Publications *mgo.Collection
	URLs         *mgo.Collection
}

func New(host string) *Mongo {
	return &Mongo{Host: host}
}

func (m *Mongo) Close() {
	m.Session.Close()
}

func (m *Mongo) Connect() (err error) {
	m.Session, err = mgo.Dial(m.Host)
	if err != nil {
		return
	}
	m.C = collections{
		Articles:     m.Session.DB(m.Prefix + ArticleCollection).C(ArticleCollection),
		Feeds:        m.Session.DB(m.Prefix + FeedCollection).C(FeedCollection),
		Keywords:     m.Session.DB(m.Prefix + KeywordCollection).C(KeywordCollection),
		Publications: m.Session.DB(m.Prefix + PublicationCollection).C(PublicationCollection),
		URLs:         m.Session.DB(m.Prefix + URLsCollection).C(URLsCollection),
	}
	return
}
