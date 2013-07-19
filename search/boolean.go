package search

import (
	"log"
	"strings"
)

type Boolean struct {
	Tree [][]*Phrase
}

func NewBoolean(query string) (b *Boolean) {
	ors := strings.Split(query, " OR ")
	b = &Boolean{
		Tree: make([][]*Phrase, len(ors)),
	}
	// Build ze tree
	for i, or := range ors {
		ands := strings.Split(or, " AND ")
		b.Tree[i] = make([]*Phrase, len(ands))
		for j, and := range ands {
			b.Tree[i][j] = NewPhrase(and)
		}
	}
	return
}

func (s *Boolean) Match(b []byte) (matches bool) {
	for i := range s.Tree {
		matches = true
		for j := range s.Tree[i] {
			if !s.Tree[i][j].Insensitive(b) {
				matches = false
				break
			}
		}
		if matches {
			log.Printf("Matched `%+q` with `%+q`", s.Tree, s.Tree[i])
			return
		}
	}
	return
}
