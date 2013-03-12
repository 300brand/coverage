package lexer

import (
	"testing"
)

// var benchmarkString comes from normalize_test.go

func BenchmarkGetWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetWords(benchmarkString)
	}
}

func TestWordStemIndex(t *testing.T) {
	/*
		s := []byte(`
			As open government gains favor and reaps benefits for the federal
			workforce and citizens, Data.gov continues to build a Google-style
			universe based on the concept and expand its mission globally.
		`)
		for _, w := range GetWords(s) {
			t.Logf("%2d %-14s %s", w.Index, w.Word, w.Stem)
		}
	*/
	t.Error("TODO")
}
