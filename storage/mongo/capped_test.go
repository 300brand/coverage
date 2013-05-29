package mongo

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestCappedCreate(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	if _, err := m.CappedCollection("TestCapped", 1); err != nil {
		t.Fatal(err)
	}

	// Ensure no error is thrown for "duplicate collection"
	if _, err := m.CappedCollection("TestCapped", 1); err != nil {
		t.Fatal(err)
	}
}

func TestCappedInsert(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	c, err := m.CappedCollection("TestCapped", 10)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err := c.Insert(bson.NewObjectId(), 0); err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < 10; i++ {
		id, err := c.NextId()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(id)
	}
	if _, err := c.NextId(); err == nil {
		t.Fatal("Expected error")
	}
}
