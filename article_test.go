package coverage

import (
	"github.com/300brand/coverage/merger"
	"testing"
)

func TestChangelog(t *testing.T) {
	a := NewArticle()
	a.Title = "Test Title"
	a.Modified("Title")

	b := NewArticle()
	b.Text.Body = Body{
		Text: []byte("Body text"),
	}
	b.Modified("Text")

	merger.Merge(a, b)

	if string(a.Text.Body.Text) != string(b.Text.Body.Text) {
		t.Error("Text struct did not merge")
	}
}
