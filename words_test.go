package coverage

import (
	"strings"
	"testing"
)

func TestWordsAll(t *testing.T) {
	words := Words{}
	a := []string{"one", "two", "one", "two", "three", "four"}
	for _, s := range a {
		words.Add(Word{Word: s})
	}
	if strings.Join(a, " ") != strings.Join(words.All, " ") {
		t.Error("All did not match")
	}
}

func TestWordsUnique(t *testing.T) {
	words := Words{}
	a := []string{"one", "two", "one", "two", "three", "four"}
	for _, s := range a {
		words.Add(Word{Word: s})
	}
	if u := strings.Join(words.Keywords, " "); u != "four one three two" {
		t.Error("Keywords did not match")
		t.Errorf("Got: %s", u)
	}
}
