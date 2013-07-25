package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"testing"
	"time"
)

func TestKeywords(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	words := []string{"a", "b", "c", "d"}

	a := coverage.NewArticle()
	a.Text.Words.Keywords = words
	a.Published = time.Now()

	if err := m.AddKeywords(a); err != nil {
		t.Fatal(err)
	}

	query := bson.M{"_id.keyword": bson.M{"$in": words}}
	n, err := m.db.C(KeywordCollection).Find(query).Count()
	if err != nil {
		t.Fatal(err)
	}
	if n != len(words) {
		t.Errorf("Invalid number of records returned: %d", n)
	}
}
