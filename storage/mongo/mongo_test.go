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
	m := New(dbURL, dbName)
	if err := m.Connect(); err != nil {
		t.Error(err)
		return
	}
	defer m.Close()

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

	m.db.DropDatabase()
}

func TestGridFSSave(t *testing.T) {
	m := New(dbURL, dbName)
	if err := m.Connect(); err != nil {
		t.Error(err)
		return
	}
	defer m.Close()
	a := coverage.NewArticle()
	a.HTML = []byte("<!DOCTYPE html><html><body><p>Test</p></body></html>")
	a.Body = coverage.Body{
		HTML: []byte("<p>Test</p>"),
		Text: []byte("Test"),
	}
	if err := m.UpdateArticle(a); err != nil {
		t.Error(err)
	}
	m.db.DropDatabase()
}
