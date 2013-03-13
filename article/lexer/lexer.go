package lexer

import (
	"bytes"
	"git.300brand.com/coverage"
	"github.com/rookii/paicehusk"
)

var StemmingEnabled = false

func GetWords(b []byte) (ws coverage.Words) {
	n := Normalize(b)
	for i, f := range bytes.Fields(n) {
		s := string(f)
		w := coverage.Word{
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
