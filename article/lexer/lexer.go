package lexer

import (
	"bytes"
	"git.300brand.com/coverage"
	"github.com/rookii/paicehusk"
)

var StemmingEnabled = false

func GetWords(b []byte) (ws coverage.Words) {
	n := Normalize(b)
	fields := bytes.Fields(n)
	for _, f := range fields {
		w := coverage.Word{
			Word:  string(f),
			Lower: string(bytes.ToLower(f)),
		}
		// Stemming adds 10x the time to split words up
		if StemmingEnabled {
			w.Stem = paicehusk.DefaultRules.Stem(w.Word)
		}
		ws.Add(w)
	}
	return
}
