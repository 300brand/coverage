package search

import (
	"bytes"
)

type Phrase struct {
	Phrase []byte
	Lower  []byte
}

func NewPhrase(s string) *Phrase {
	p := &Phrase{
		Phrase: []byte(s),
	}
	p.Lower = bytes.ToLower(p.Phrase)
	return p
}

func (p *Phrase) Insensitive(b []byte) bool {
	return bytes.Contains(bytes.ToLower(b), p.Lower)
}

func (p *Phrase) Match(b []byte) bool {
	return bytes.Contains(b, p.Phrase)
}

func (p *Phrase) String() string {
	return string(p.Phrase)
}
