package service

import (
	"git.300brand.com/coverage"
)

type ReportService interface {
	Update(*coverage.Report) error
}
