package search

import (
	"github.com/300brand/coverage/article/lexer"
	"github.com/300brand/logger"
	"strings"
)

type Boolean struct {
	Query   string
	Tree    [][]*Phrase
	Exclude []*Phrase
}

func NewBoolean(query string) (b *Boolean) {
	b = &Boolean{
		Query: query,
	}
	if qNot := strings.SplitN(query, " NOT ", 2); len(qNot) == 2 {
		query = qNot[0]
		nots := strings.Split(qNot[1], " OR ")
		b.Exclude = make([]*Phrase, len(nots))
		for i := range nots {
			b.Exclude[i] = NewPhrase(nots[i])
		}
	}
	ors := strings.Split(query, " OR ")
	b.Tree = make([][]*Phrase, len(ors))
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
	for i := range s.Exclude {
		if s.Exclude[i].Insensitive(b) {
			logger.Trace.Printf("Matched: `%+q` `NOT %s`", s.Exclude, s.Exclude[i])
			return false
		}
	}
	for i := range s.Tree {
		matches = true
		for j := range s.Tree[i] {
			if !s.Tree[i][j].Insensitive(b) {
				matches = false
				break
			}
		}
		if matches {
			logger.Trace.Printf("Matched `%+q` with `%+q`", s.Tree, s.Tree[i])
			break
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
