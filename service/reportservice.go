package service

import (
	"github.com/300brand/coverage"
)

type ReportService interface {
	Update(*coverage.Report) error
}
