package lexer

import (
	"bytes"
	"github.com/rookii/paicehusk"
)

func GetWords(b []byte) (w Words) {
	n := Normalize(b)
	for i, f := range bytes.Fields(n) {
		s := string(f)
		w.Add(Word{
			Word:  s,
			Stem:  paicehusk.DefaultRules.Stem(s),
			Index: i,
		})
	}
	return
}
