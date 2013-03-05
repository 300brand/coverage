package lexer

import (
	"testing"
)

func TestWordStemIndex(t *testing.T) {
	s := []byte(`
		As open government gains favor and reaps benefits for the federal
		workforce and citizens, Data.gov continues to build a Google-style
		universe based on the concept and expand its mission globally.
	`)
	for _, w := range GetWords(s) {
		t.Logf("%2d %-14s %s", w.Index, w.Word, w.Stem)
	}
}
