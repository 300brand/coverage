package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"github.com/jbaikge/logger"
	"labix.org/v2/mgo/bson"
	"time"
)

type PublicationService struct {
	m *Mongo
}

const PublicationCollection = "Publications"

var _ service.PublicationService = &PublicationService{}

func NewPublicationService(m *Mongo) *PublicationService {
	return &PublicationService{m: m}
}

func (s *PublicationService) Update(p *coverage.Publication) error {
	p.Log.Service("mongo.PublicationService")
	return s.m.UpdatePublication(p)
}

func (m *Mongo) GetPublication(query interface{}, p *coverage.Publication) (err error) {
	logger.Trace.Printf("GetPublication: called %+v", query)
	switch v := query.(type) {
	case bson.ObjectId:
		query = bson.M{"_id": v}
	}
	err = m.C.Publications.Find(query).One(p)
	return
}

func (m *Mongo) GetPublications(query interface{}, sort string, skip, limit int, p *[]*coverage.Publication) (err error) {
	logger.Trace.Printf("GetPublications: called")
	logger.Trace.Printf("query: %+v", query)
	logger.Trace.Printf("sort: %s skip: %d limit: %d", sort, skip, limit)
	q := m.C.Publications.Find(query)
	if sort != "" {
		q.Sort(sort)
	}
	if skip > 0 {
		q.Skip(skip)
	}
	if limit > 0 {
		q.Limit(limit)
	}
	return q.All(p)
}

func (m *Mongo) UpdatePublication(p *coverage.Publication) (err error) {
	logger.Trace.Printf("UpdatePublication: called")
	p.Updated = time.Now()
	_, err = m.C.Publications.UpsertId(p.ID, p)
	return
}

func (m *Mongo) PublicationIncFeeds(id bson.ObjectId, delta int) (err error) {
	logger.Trace.Printf("PublicationIncFeeds: called %s %+d", id.Hex(), delta)
	return m.C.Publications.UpdateId(id, bson.M{
		"$inc": bson.M{
			"numfeeds": delta,
		},
		"$set": bson.M{
			"updated": time.Now(),
		},
	})
}

func (m *Mongo) PublicationIncArticles(id bson.ObjectId, delta int) (err error) {
	logger.Trace.Printf("PublicationIncArticles: called %s %+d", id.Hex(), delta)
	return m.C.Publications.UpdateId(id, bson.M{
		"$inc": bson.M{
			"numarticles": delta,
		},
		"$set": bson.M{
			"updated": time.Now(),
		},
	})
}
