package lexer

import (
	"testing"
)

// var benchmarkString comes from normalize_test.go

func BenchmarkWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Words(benchmarkString)
	}
}

func BenchmarkStemming(b *testing.B) {
	StemmingEnabled = true
	for i := 0; i < b.N; i++ {
		Words(benchmarkString)
	}
	StemmingEnabled = false
}

func TestWordsLen(t *testing.T) {
	s := []byte(`
		As open government gains favor and reaps benefits for the federal
		workforce and citizens, Data.gov continues to build a Google-style
		universe based on the concept and expand its mission globally.
	`)
	words := Words(s)
	if l := len(words); l != 31 {
		t.Errorf("Got %d words", l)
	}
}
