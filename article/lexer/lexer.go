package lexer

import (
	"bytes"
)

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

func fieldFilter(r rune) bool {
	for _, f := range filterFuncs {
		if f(r) {
			return true
		}
	}
	return false
}
