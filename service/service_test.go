package service

import (
	"git.300brand.com/coverage"
	"testing"
)

type as struct{}
type fs struct{}
type ps struct{}
type rs struct{}

func (s as) Update(a *coverage.Article) error     { return nil }
func (s fs) Update(f *coverage.Feed) error        { return nil }
func (s ps) Update(p *coverage.Publication) error { return nil }
func (s rs) Update(r *coverage.Report) error      { return nil }

func TestArticleService(t *testing.T)     { var _ ArticleService = as{} }
func TestFeedService(t *testing.T)        { var _ FeedService = fs{} }
func TestPublicationService(t *testing.T) { var _ PublicationService = ps{} }
func TestReportService(t *testing.T)      { var _ ReportService = rs{} }
