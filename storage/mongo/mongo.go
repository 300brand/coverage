package mongo

import (
	"github.com/300brand/logger"
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
	ArticleQ      *mgo.Collection
	FeedQ         *mgo.Collection
	Feeds         *mgo.Collection
	GroupSearch   *mgo.Collection
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
	m.C = collections{
		Articles:     m.Session.DB(m.Prefix + ArticleCollection).C(ArticleCollection),
		ArticleQ:     m.Session.DB(m.Prefix + ArticleCollection).C(ArticleQueueCollection),
		FeedQ:        m.Session.DB(m.Prefix + FeedCollection).C(FeedQueueCollection),
		Feeds:        m.Session.DB(m.Prefix + FeedCollection).C(FeedCollection),
		Keywords:     m.Session.DB(m.Prefix + KeywordCollection).C(KeywordCollection),
		Publications: m.Session.DB(m.Prefix + PublicationCollection).C(PublicationCollection),
		Search:       m.Session.DB(m.Prefix + SearchCollection).C(SearchCollection),
		GroupSearch:  m.Session.DB(m.Prefix + SearchCollection).C(GroupSearchCollection),
		URLs:         m.Session.DB(m.Prefix + URLsCollection).C(URLsCollection),
		session:      m.Session,
	}
	return
}

func (m *Mongo) Copy() collections {
	newSession := m.Session.Copy()
	c := collections{
		Articles:     newSession.DB(m.Prefix + ArticleCollection).C(ArticleCollection),
		ArticleQ:     newSession.DB(m.Prefix + ArticleCollection).C(ArticleQueueCollection),
		FeedQ:        newSession.DB(m.Prefix + FeedCollection).C(FeedQueueCollection),
		Feeds:        newSession.DB(m.Prefix + FeedCollection).C(FeedCollection),
		Keywords:     newSession.DB(m.Prefix + KeywordCollection).C(KeywordCollection),
		Publications: newSession.DB(m.Prefix + PublicationCollection).C(PublicationCollection),
		Search:       newSession.DB(m.Prefix + SearchCollection).C(SearchCollection),
		GroupSearch:  newSession.DB(m.Prefix + SearchCollection).C(GroupSearchCollection),
		URLs:         newSession.DB(m.Prefix + URLsCollection).C(URLsCollection),
		session:      newSession,
	}
	return c
}

// NOOP - Copy -> Close creates too many open connections to mongo during
// searches.
func (c collections) Close() {
	c.session.Close()
}
