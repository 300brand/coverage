package coverage

import (
	"git.300brand.com/coverage/merger"
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
	b.Modified("Body")

	merger.Merge(a, b)

	if string(a.Text.Body.Text) != string(b.Text.Body.Text) {
		t.Error("Body struct did not merge")
	}
}
