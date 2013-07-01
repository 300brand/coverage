package search

import (
	"git.300brand.com/coverage"
)

type Filter interface {
	Matches(*coverage.Article, []string, *Stats) error
}

type Stats struct {
	All bool
}
