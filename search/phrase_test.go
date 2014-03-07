package search

import (
	"testing"
)

func TestPhrase(t *testing.T) {
	tests := []struct {
		Haystack    []byte
		Needle      string
		Match       bool
		Insensitive bool
	}{
		{
			[]byte("monkey poo"),
			"monkey poo",
			true,
			true,
		},
		{
			[]byte("throughput"),
			"HP",
			false,
			false,
		},
		{
			[]byte("fast network with high HP throughput"),
			"HP",
			true,
			true,
		},
		{
			[]byte("throughput in the HP offering"),
			"HP",
			true,
			true,
		},
		{
			[]byte("said HP to a journalist"),
			"HP",
			true,
			true,
		},
		{
			[]byte("HP's greatest invention"),
			"HP",
			true,
			true,
		},
		{
			[]byte("said HP's representative"),
			"HP",
			true,
			true,
		},
		{
			[]byte("washington, march 6 "),
			"SHI",
			false,
			false,
		},
	}
	for i, test := range tests {
		p := NewPhrase(test.Needle)
		if p.Match(test.Haystack) != test.Match {
			t.Errorf("[%d] Match failed", i)
		}
		if p.Insensitive(test.Haystack) != test.Insensitive {
			t.Errorf("[%d] Insensitive failed", i)
		}
	}
}

func BenchmarkPhraseInsensitive(b *testing.B) {
	p := NewPhrase("HP")
	haystack := []byte(`throughput in the HP offering`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Insensitive(haystack)
	}
}
