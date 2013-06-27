package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestKeywords(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	insertKeywordArticles(t, m)

	info, err := m.ReduceKeywords(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)
}

func TestKeywordArticleIds(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	kws := []string{"a", "f", "h"}
	articles := insertKeywordArticles(t, m)
	m.ReduceKeywords(nil)
	ids, err := m.KeywordArticleIds(kws, time.Time{}, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ids); l != 4 {
		t.Errorf("Invalid number of ids returned: %d", l)
	}
	for _, id := range ids {
		a := articles[id]
		found := false
		for _, kw := range kws {
			for _, w := range a.Text.Words.Keywords {
				if w == kw {
					found = true
				}
			}
		}
		if !found {
			t.Errorf("Did not find %v in %s with %v", kws, id, a.Text.Words.Keywords)
		}
	}
}

func insertKeywordArticles(t *testing.T, m *Mongo) (articles map[bson.ObjectId]*coverage.Article) {
	words := strings.Fields("a b c d e f g h")
	max := len(words) - 2
	articles = make(map[bson.ObjectId]*coverage.Article, max)
	for i := 0; i < max; i++ {
		a := coverage.NewArticle()
		a.URL, _ = url.Parse("http://" + a.ID.Hex())
		a.Published = time.Now().Add(time.Duration(-1*(max-i)) * time.Hour)
		a.Text.Words.Keywords = words[i : i+3]
		if err := m.UpdateArticle(a); err != nil {
			t.Fatal(err)
		}
		articles[a.ID] = a
	}
	return
}
