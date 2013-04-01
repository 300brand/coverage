package lexer

import (
	"strings"
	"testing"
)

func TestKeywords(t *testing.T) {
	in := []byte("This is my test string. It's just a small one, but there is potential.")
	out := strings.Fields("just one potential small string test")
	keywords := Keywords(in)

	if len(out) != len(keywords) {
		t.Fatalf("Lengths do not match - len(out): %d; len(keywords): %d", len(out), len(keywords))
	}

	for i, kw := range keywords {
		if o := out[i]; o != kw {
			t.Errorf("Expected: %s; Got: %s", o, kw)
		}
	}
}
