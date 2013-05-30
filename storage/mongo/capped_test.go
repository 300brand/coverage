package mongo

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestCappedAll(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	c, err := m.CappedIdCollection("TestCapped", 10)
	if err != nil {
		t.Fatal(err)
	}

	all := make([]bson.ObjectId, 10)
	for i := range all {
		all[i] = bson.NewObjectId()
		if err := c.Insert(all[i], 0); err != nil {
			t.Fatal(err)
		}
	}

	ids, err := c.All()
	if err != nil {
		t.Fatal(err)
	}
	if aLen, iLen := len(all), len(ids); aLen != iLen {
		t.Fatalf("All Length mismatch: %d != %d", aLen, iLen)
	}
	for i := range ids {
		if aLen, iLen := all[i], ids[i]; aLen != iLen {
			t.Errorf("All Mismatch: %s != %s", aLen, iLen)
		}
	}

	// Test FIFO ordering with Next usage
	for _, id := range all[:5] {
		next, err := c.Next()
		if err != nil {
			t.Fatal(err)
		}
		if next != id {
			t.Errorf("Next Mismatch: %s != %s", id, next)
		}
	}

	// Ensure the result of All returns the last half of unprocessed
	ids, err = c.All()
	if err != nil {
		t.Fatal(err)
	}
	if aLen, iLen := len(all), len(ids); aLen != iLen {
		t.Fatalf("Half Length mismatch: %d != %d", aLen, iLen)
	}
	for i := range ids {
		if a, id := all[5+i], ids[i]; a != id {
			t.Errorf("Remaining mismatch: %s != %s", a, id)
		}
	}
}

func TestCappedCreate(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	if _, err := m.CappedIdCollection("TestCapped", 1); err != nil {
		t.Fatal(err)
	}

	// Ensure no error is thrown for "duplicate collection"
	if _, err := m.CappedIdCollection("TestCapped", 1); err != nil {
		t.Fatal(err)
	}
}

func TestCappedInsert(t *testing.T) {
	m := connect(t)
	defer cleanup(m)

	c, err := m.CappedIdCollection("TestCapped", 10)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err := c.Insert(bson.NewObjectId(), 0); err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < 10; i++ {
		id, err := c.Next()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(id)
	}
	if _, err := c.Next(); err == nil {
		t.Fatal("Expected error")
	}
}
