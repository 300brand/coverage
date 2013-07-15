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
