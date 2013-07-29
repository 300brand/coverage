package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
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

func (m *Mongo) GetPublication(query interface{}) (p *coverage.Publication, err error) {
	p = &coverage.Publication{}
	err = m.C.Publications.Find(query).One(p)
	return
}

func (m *Mongo) UpdatePublication(p *coverage.Publication) (err error) {
	p.Updated = time.Now()
	_, err = m.C.Publications.UpsertId(p.ID, p)
	return
}
