package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"labix.org/v2/mgo"
)

type PublicationService struct {
	m *Mongo
}

const PublicationCollection = "Publications"

var _ service.PublicationService = &PublicationService{}

func init() {
	indexes[PublicationCollection] = []mgo.Index{
		mgo.Index{
			Key:        []string{"url"},
			Background: true,
			DropDups:   true,
			Sparse:     false,
			Unique:     true,
		},
	}
}

func NewPublicationService(m *Mongo) *PublicationService {
	return &PublicationService{m: m}
}

func (s *PublicationService) Update(p *coverage.Publication) error {
	p.Log.Service("mongo.PublicationService")
	return s.m.UpdatePublication(p)
}

func (m *Mongo) GetPublication(query interface{}) (p *coverage.Publication, err error) {
	p = &coverage.Publication{}
	err = m.db.C(PublicationCollection).Find(query).One(p)
	return
}

func (m *Mongo) UpdatePublication(p *coverage.Publication) (err error) {
	p.Log.Debug("mongo.UpdatePublication")
	_, err = m.db.C(PublicationCollection).UpsertId(p.ID, p)
	return
}
