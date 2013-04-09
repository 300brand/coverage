package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"labix.org/v2/mgo"
	"time"
)

type ReportService struct {
	m *Mongo
}

const ReportCollection = "Reports"

var _ service.ReportService = &ReportService{}

func init() {
	indexes[ReportCollection] = []mgo.Index{}
}

func NewReportService(m *Mongo) *ReportService {
	return &ReportService{m: m}
}

func (s *ReportService) Update(r *coverage.Report) error {
	r.Log.Service("mongo.ReportService")
	return s.m.UpdateReport(r)
}

func (m *Mongo) UpdateReport(r *coverage.Report) (err error) {
	if err = m.EnsureIndexSet(ReportCollection, indexes[ReportCollection]); err != nil {
		return
	}

	if r.Added.Equal(time.Time{}) {
		r.Added = time.Now()
	}

	_, err = m.db.C(ReportCollection).UpsertId(r.ID, r)
	return
}
