package mongo

import (
	"github.com/300brand/coverage"
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestStats(t *testing.T) {
	m := connect(t)
	defer cleanup(m)
	var err error

	err = m.C.Publications.Insert(
		&coverage.Publication{
			ID:          bson.NewObjectId(),
			NumArticles: 3,
		},
		&coverage.Publication{
			ID:      bson.NewObjectId(),
			Deleted: true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	err = m.C.Feeds.Insert(
		&coverage.Feed{
			ID:            bson.NewObjectId(),
			PublicationId: bson.NewObjectId(),
		},
		&coverage.Feed{
			ID:            bson.NewObjectId(),
			PublicationId: bson.NewObjectId(),
			Deleted:       true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	err = m.C.Articles.Insert(
		&coverage.Article{
			ID:            bson.NewObjectId(),
			FeedId:        bson.NewObjectId(),
			PublicationId: bson.NewObjectId(),
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	err = m.C.FeedQ.Insert(struct{ FeedId bson.ObjectId }{bson.NewObjectId()})
	if err != nil {
		t.Fatal(err)
	}

	s := new(Stats)
	if err := m.GetStats(s); err != nil {
		t.Fatal(err)
	}
	if s.Hash() != "1.1.2.1.1.2.3.1.1" {
		t.Errorf("Invalid hash: %s", s.Hash())
	}
}
