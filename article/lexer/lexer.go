package lexer

import (
	"bytes"
	"github.com/rookii/paicehusk"
)

var StemmingEnabled = false

func GetWords(b []byte) (ws Words) {
	n := Normalize(b)
	for i, f := range bytes.Fields(n) {
		s := string(f)
		w := Word{
			Word:  s,
			Index: i,
		}
		// Stemming adds 10x the time to split words up
		if StemmingEnabled {
			w.Stem = paicehusk.DefaultRules.Stem(s)
		}
		ws.Add(w)
	}
	return
}
