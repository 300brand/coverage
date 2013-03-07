package service

import (
	"git.300brand.com/coverage"
)

type Service interface {
	Update(*coverage.Article) error
}
