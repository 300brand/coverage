package search

import (
	"git.300brand.com/coverage"
	"testing"
)

var all All

func BenchmarkAll(b *testing.B) {
	b.ReportAllocs()
	var (
		article = coverage.NewArticle()
		stats   Stats
		terms   = []string{"bar", "foo"}
	)
	article.Text.Words.Keywords = []string{"bar", "baz", "biz", "fee", "foo"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		all.Matches(article, terms, &stats)
	}
}

func TestAllFail(t *testing.T) {
	var (
		article = coverage.NewArticle()
		stats   Stats
		terms   = []string{"boo", "far"}
	)
	article.Text.Words.Keywords = []string{"bar", "baz", "biz", "fee", "foo"}
	if err := all.Matches(article, terms, &stats); err != nil {
		t.Fatal(err)
	}
	if stats.All {
		t.Fatal("Matched")
	}
}

func TestAllPass(t *testing.T) {
	var (
		article = coverage.NewArticle()
		stats   Stats
		terms   = []string{"bar", "foo"}
	)
	article.Text.Words.Keywords = []string{"bar", "baz", "biz", "fee", "foo"}
	if err := all.Matches(article, terms, &stats); err != nil {
		t.Fatal(err)
	}
	if !stats.All {
		t.Fatal("Did not match")
	}
}
