package body

import (
	"github.com/300brand/coverage"
	"testing"
)

func TestArticleExtract(t *testing.T) {
	a := coverage.NewArticle()
	a.Text.HTML = []byte(`<!DOCTYPE html><head></head><body><div><p>First page body</p><p>More first page content.</p></div></body></html>`)
	a.Text.Pages = [][]byte{
		[]byte(`<!DOCTYPE html><head></head><body><div><p>Second page body</p><p>TMore second page content.</p></div></body></html>`),
		[]byte(`<!DOCTYPE html><head></head><body><div><p>Third page body</p><p>More third page content.</p></div></body></html>`),
	}

	if err := SetBody(a); err != nil {
		t.Errorf("Error setting body: %s", err)
	}

	t.Logf("Multipage HTML: %s", a.Text.Body.HTML)
	t.Logf("Multipage Text: %s", a.Text.Body.Text)
}
