package service

import (
	"git.300brand.com/coverage"
)

type ArticleService interface {
	Update(*coverage.Article) error
}
