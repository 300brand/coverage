package lexer

import (
	"bytes"
)

type Word struct {
	Word  string
	Stem  string
	Index int
}

type Words []Word

func GetWords(b []byte) (w Words) {
	fields := bytes.FieldsFunc(b, fieldFilter)
	for i, f := range fields {
		w.Add(Word{
			Word:  string(f),
			Index: i,
		})
	}
	return
}

func (w Word) String() string {
	return w.Word
}

func (ws *Words) Add(w Word) {
	*ws = append(*ws, w)
}

func (ws *Words) Strings() (s []string) {
	for _, w := range *ws {
		s = append(s, w.Word)
	}
	return
}

func fieldFilter(r rune) bool {
	for _, f := range filterFuncs {
		if f(r) {
			return true
		}
	}
	return false
}
