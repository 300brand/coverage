package service

import (
	"github.com/300brand/coverage"
)

type ArticleService interface {
	Update(*coverage.Article) error
}
