package mongo

import (
	"github.com/300brand/coverage"
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestArticleQueue(t *testing.T) {
	m := connect(t)
	//defer close(m)

	var in, out coverage.Article

	in = *coverage.NewArticle()
	in.FeedId = bson.NewObjectId()
	in.PublicationId = bson.NewObjectId()
	if err := m.ArticleQueueAdd(&in); err != nil {
		t.Fatalf("Error enqueuing: %s", err)
	}

	if err := m.ArticleQueueNext(&out); err != nil {
		t.Fatalf("Error fetching next article: %s", err)
	}

	if err := m.ArticleQueueNext(&out); err == nil {
		t.Fatalf("Expected error when fetching from queue")
	}

	if err := m.ArticleQueueRemove(in.ID); err != nil {
		t.Fatalf("Error removing article from queue: %s", err)
	}
}
