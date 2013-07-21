package search

import (
	"git.300brand.com/coverage/article/lexer"
	"log"
	"strings"
)

type Boolean struct {
	Query string
	Tree  [][]*Phrase
}

func NewBoolean(query string) (b *Boolean) {
	ors := strings.Split(query, " OR ")
	b = &Boolean{
		Query: query,
		Tree:  make([][]*Phrase, len(ors)),
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

func (s *Boolean) MinTerms() (min int) {
	min = 255
	for i := range s.Tree {
		n := 0
		for _, p := range s.Tree[i] {
			n += len(lexer.Keywords(p.Lower))
		}
		if n < min {
			min = n
		}
	}
	return
}
