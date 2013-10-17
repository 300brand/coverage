package mongo

import (
	"github.com/jbaikge/logger"
	"labix.org/v2/mgo"
)

type Mongo struct {
	C       collections
	Host    string
	Prefix  string
	Session *mgo.Session
}

type collections struct {
	Articles      *mgo.Collection
	FeedQ         *mgo.Collection
	Feeds         *mgo.Collection
	Keywords      *mgo.Collection
	Publications  *mgo.Collection
	Search        *mgo.Collection
	SearchResults *mgo.Collection
	URLs          *mgo.Collection
	session       *mgo.Session
}

func New(host string) *Mongo {
	return &Mongo{Host: host}
}

func (m *Mongo) Close() {
	logger.Trace.Printf("Close: called")
	m.Session.Close()
}

func (m *Mongo) Connect() (err error) {
	logger.Trace.Printf("Connecting to %s", m.Host)
	m.Session, err = mgo.Dial(m.Host)
	if err != nil {
		return
	}
	m.C = m.Copy()
	return
}

func (m *Mongo) Copy() collections {
	s := m.Session.Copy()
	return collections{
		Articles:     s.DB(m.Prefix + ArticleCollection).C(ArticleCollection),
		FeedQ:        s.DB(m.Prefix + FeedCollection).C(FeedQueueCollection),
		Feeds:        s.DB(m.Prefix + FeedCollection).C(FeedCollection),
		Keywords:     s.DB(m.Prefix + KeywordCollection).C(KeywordCollection),
		Publications: s.DB(m.Prefix + PublicationCollection).C(PublicationCollection),
		Search:       s.DB(m.Prefix + SearchCollection).C(SearchCollection),
		URLs:         s.DB(m.Prefix + URLsCollection).C(URLsCollection),
		session:      s,
	}
}

func (c collections) Close() {
	c.session.Close()
}
