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
	// Operate on slices of the whole and regain the index of the phrase at
	// every loop
	for i := bytes.Index(s, p.Lower); i > -1 && i < len(s); s, i = s[i+len(p.Lower):], bytes.Index(s[i+len(p.Lower):], p.Lower) {
		// Beginning of string automatically means there's nothing preceeding;
		// Check for alphanumerics and reslice if found
		if i > 0 {
			preRune := rune(s[i-1])
			if unicode.IsLetter(preRune) || unicode.IsNumber(preRune) {
				continue
			}
		}
		// End of string automatically means there's nothing following; Check
		// for alphanumerics and reslice if found
		if i+len(p.Lower) < len(b) {
			postRune := rune(s[i+len(p.Lower)])
			if unicode.IsLetter(postRune) || unicode.IsNumber(postRune) {
				continue
			}
		}
		return true
	}
	return false
}

func (p *Phrase) Match(b []byte) bool {
	return bytes.Contains(b, p.Phrase)
}

func (p *Phrase) String() string {
	return string(p.Phrase)
}
