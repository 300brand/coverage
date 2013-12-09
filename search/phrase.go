package search

import (
	"bytes"
	"unicode"
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
	s := bytes.ToLower(b)
	i := bytes.Index(s, p.Lower)
	for i > -1 && i < len(s) {
		if i == 0 || unicode.IsSpace(rune(s[i-1])) {
			return true
		}
		// determine next slice origin
		si := i + len(p.Lower) + 1
		if si >= len(s) {
			return false
		}
		i = bytes.Index(s[si:], p.Lower) + si
	}
	return false
}

func (p *Phrase) Match(b []byte) bool {
	return bytes.Contains(b, p.Phrase)
}

func (p *Phrase) String() string {
	return string(p.Phrase)
}
