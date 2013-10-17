package service

import (
	"github.com/300brand/coverage"
)

type PublicationService interface {
	Update(*coverage.Publication) error
}
