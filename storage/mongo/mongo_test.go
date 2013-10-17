// +build !goci

package mongo

import (
	"github.com/300brand/coverage"
	"labix.org/v2/mgo/bson"
	"net/url"
	"testing"
)

var host = "localhost"

func TestConnect(t *testing.T) {
	if err := New(host).Connect(); err != nil {
		t.Error(err)
	}
}

func TestArticleSave(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	in := coverage.NewArticle()
	in.Title = "Test Title"
	in.URL, _ = url.Parse("http://www.google.com/search?q=search")
	in.Modified()
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

func cleanup(m *Mongo) {
	m.C.Articles.Database.DropDatabase()
	m.C.Feeds.Database.DropDatabase()
	m.C.Keywords.Database.DropDatabase()
	m.C.Publications.Database.DropDatabase()
	m.C.Search.Database.DropDatabase()
	m.C.URLs.Database.DropDatabase()
	m.Close()
}

func connect(t *testing.T) (m *Mongo) {
	if testing.Short() {
		t.Skip("Short tests running")
	}
	m = New(host)
	m.Prefix = "Test"
	if err := m.Connect(); err != nil {
		t.Error(err)
		return
	}
	return
}
