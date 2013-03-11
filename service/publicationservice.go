package service

import (
	"git.300brand.com/coverage"
)

type PublicationService interface {
	Update(*coverage.Publication) error
}
