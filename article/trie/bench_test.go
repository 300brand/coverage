package trie

import (
	"testing"
)

var wordlist = []string{
	"a", "about", "above", "after", "again", "against", "all", "am", "an",
	"and", "any", "are", "aren't", "as", "at", "be", "because", "been",
	"before", "being", "below", "between", "both", "but", "by", "can't",
	"cannot", "could", "couldn't", "did", "didn't", "do", "does", "doesn't",
	"doing", "don't", "down", "during", "each", "few", "for", "from",
	"further", "had", "hadn't", "has", "hasn't", "have", "haven't", "having",
	"he", "he'd", "he'll", "he's", "her", "here", "here's", "hers", "herself",
	"him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm", "i've",
	"if", "in", "into", "is", "isn't", "it", "it's", "its", "itself", "let's",
	"me", "more", "most", "mustn't", "my", "myself", "no", "nor", "not", "of",
	"off", "on", "once", "only", "or", "other", "ought", "our", "ours",
	"ourselves", "out", "over", "own", "same", "shan't", "she", "she'd",
	"she'll", "she's", "should", "shouldn't", "so", "some", "such", "than",
	"that", "that's", "the", "their", "theirs", "them", "themselves", "then",
	"there", "there's", "these", "they", "they'd", "they'll", "they're",
	"they've", "this", "those", "through", "to", "too", "under", "until", "up",
	"very", "was", "wasn't", "we", "we'd", "we'll", "we're", "we've", "were",
	"weren't", "what", "what's", "when", "when's", "where", "where's", "which",
	"while", "who", "who's", "whom", "why", "why's", "with", "won't", "would",
	"wouldn't", "you", "you'd", "you'll", "you're", "you've", "your", "yours",
	"yourself", "yourselves",
}

func TestHasAll(t *testing.T) {
	tr := New()
	for _, w := range wordlist {
		tr.Add(w)
	}
	for _, w := range wordlist {
		if !tr.Has(w) {
			t.Errorf("Expected to find: %s", w)
		}
	}
}

func TestCant(t *testing.T) {
	tr := New()
	tr.Add("can't")
	tr.Add("cannot")
	tr.Add("could")
	t.Log(tr.Dump())
	if !tr.Has("can't") {
		t.Error("Expected to find can't")
	}
}

func BenchmarkAdd(b *testing.B) {
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Add(wordlist[i%len(wordlist)])
	}
}

func BenchmarkHas(b *testing.B) {
	b.StopTimer()
	tr := New()
	for _, w := range wordlist {
		tr.Add(w)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Has(wordlist[i%len(wordlist)])
	}
}
