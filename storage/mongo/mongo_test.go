package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"net/url"
	"testing"
	"time"
)

var (
	dbURL  = "localhost"
	dbName = "CoverageTest"
)

func TestConnect(t *testing.T) {
	if err := New(dbURL, dbName).Connect(); err != nil {
		t.Error(err)
	}
}

func TestArticleSave(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	in := coverage.NewArticle()
	in.Title = "Test Title"
	in.URL, _ = url.Parse("http://www.google.com/search?q=search")
	in.Times.Updated = time.Now()
	m.UpdateArticle(in)

	out, err := m.GetArticle(bson.M{"_id": in.ID})
	if err != nil {
		t.Error(err)
		return
	}
	if in.Title != out.Title {
		t.Error("Title Mismatch")
	}
	if in.URL.String() != out.URL.String() {
		t.Error("URL Mismatch")
		t.Logf("In:  %s", in.URL)
		t.Logf("Out: %s", out.URL)
	}
}

func TestGridFSSave(t *testing.T) {
	m := connect(t)
	defer cleanup(m)
	a := coverage.NewArticle()
	a.HTML = []byte("<!DOCTYPE html><html><body><p>Test</p></body></html>")
	a.Body = coverage.Body{
		HTML: []byte("<p>Test</p>"),
		Text: []byte("Test"),
	}
	if err := m.UpdateArticle(a); err != nil {
		t.Error(err)
	}
}

func TestGridFSUpdate(t *testing.T) {
	m := connect(t)
	defer cleanup(m)
	a := coverage.NewArticle()

	a.HTML = []byte("This is the first document")
	m.UpdateArticle(a)

	a.HTML = []byte("This is the second document")
	m.UpdateArticle(a)
}

func TestNoDupes(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	a := coverage.NewArticle()
	a.URL, _ = url.Parse("http://google.com")
	a.Title = "Google Homepage"
	if err := m.UpdateArticle(a); err != nil {
		t.Error(err)
	}

	b := coverage.NewArticle()
	b.URL, _ = url.Parse(a.URL.String())
	b.Title = "Random Redirect"
	if err := m.UpdateArticle(b); err == nil {
		t.Error("No error encountered for duplicate URL")
	}

	c := coverage.NewArticle()
	c.URL, _ = url.Parse("http://redirect.me/to/google.com")
	c.Title = b.Title
	if err := m.UpdateArticle(c); err != nil {
		t.Error(err)
	}
}

func cleanup(m *Mongo) {
	m.Close()
}

func connect(t *testing.T) (m *Mongo) {
	m = New(dbURL, dbName)
	if err := m.Connect(); err != nil {
		t.Error(err)
		return
	}
	m.db.DropDatabase()
	return
}
