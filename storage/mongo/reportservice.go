package mongo

import (
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
)

type ReportService struct {
	m *Mongo
}

const ReportCollection = "Reports"

var _ service.ReportService = &ReportService{}

func NewReportService(m *Mongo) *ReportService {
	return &ReportService{m: m}
}

func (s *ReportService) Update(r *coverage.Report) error {
	r.Log.Service("mongo.ReportService")
	return s.m.UpdateReport(r)
}

func (m *Mongo) UpdateReport(r *coverage.Report) (err error) {
	if len(r.Feeds) > 0 {
		r.FeedIds, err = m.FeedIds(r.Feeds)
		if err != nil {
			return err
		}

	}

	_, err = m.Session.DB(ReportCollection).C(ReportCollection).UpsertId(r.ID, r)
	return
}
