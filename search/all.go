package search

import (
	"fmt"
	"git.300brand.com/coverage"
	"sort"
)

type All struct{}

var _ Filter = All{}

func (f All) Matches(a *coverage.Article, terms []string, s *Stats) (err error) {
	s.All = false
	if len(terms) == 0 {
		return fmt.Errorf("Empty terms")
	}
	if len(a.Text.Words.Keywords) == 0 {
		return fmt.Errorf("Empty keywords")
	}
	if !sort.StringsAreSorted(a.Text.Words.Keywords) {
		return fmt.Errorf("Keywords are not sorted")
	}
	for _, term := range terms {
		if i := sort.SearchStrings(a.Text.Words.Keywords, term); a.Text.Words.Keywords[i] != term {
			return
		}
	}
	s.All = true
	return
}
