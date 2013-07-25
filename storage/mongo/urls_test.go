package mongo

import (
	"labix.org/v2/mgo/bson"
	"net/url"
	"testing"
)

func TestAddURL(t *testing.T) {
	m := connect(t)
	//defer cleanup(m)

	id1, id2 := bson.NewObjectId(), bson.NewObjectId()
	url1, _ := url.Parse("http://www.google.com")
	url2, _ := url.Parse("http://www.google.com/search")

	if err := m.AddURL(url1, id1); err != nil {
		t.Fatal(err)
	}
	if err := m.AddURL(url2, id2); err != nil {
		t.Fatal(err)
	}
	if err := m.AddURL(url1, id2); err == nil {
		t.Fatalf("Expected error with: {%s, %s}", url1, id2)
	}
}
