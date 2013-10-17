package search

import (
	"github.com/300brand/coverage"
)

type Filter interface {
	Matches(*coverage.Article, []string, *Stats) error
}

type Stats struct {
	All bool
}
