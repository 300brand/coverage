package idqueue

import (
	"labix.org/v2/mgo/bson"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	q := connect(t, "TestConnect")
	defer close(q)
}

func TestGetBlank(t *testing.T) {
	q := connect(t, "TestGetBlank")
	defer close(q)
	ids, err := q.Get()
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 0 {
		t.Error("Expected id length of 0")
	}
}

func TestPut(t *testing.T) {
	q := connect(t, "TestPut")
	defer close(q)
	ids := []bson.ObjectId{
		bson.NewObjectId(),
	}
	if err := q.Put(ids); err != nil {
		t.Fatal(err)
	}
}

func TestPutGet(t *testing.T) {
	q := connect(t, "TestPutGet")
	defer close(q)
	ids := []bson.ObjectId{
		bson.NewObjectId(),
	}
	if err := q.Put(ids); err != nil {
		t.Fatal(err)
	}
	get, err := q.Get()
	if err != nil {
		t.Fatal(err)
	}
	if gLen, iLen := len(get), len(ids); gLen != iLen {
		t.Fatal("Length mismatch: %d != %d", gLen, iLen)
	}
	if g, i := get[0], ids[0]; g != i {
		t.Fatal("ID Mismatch")
	}
}

func TestPush(t *testing.T) {
	q := connect(t, "TestPush")
	defer close(q)
	// Reset/init ids
	if err := q.Put([]bson.ObjectId{}); err != nil {
		t.Fatal(err)
	}
	all := make([]bson.ObjectId, 10)
	for i := 0; i < 10; i++ {
		all[i] = bson.NewObjectId()
	}

	// Push
	for i := range all {
		if err := q.Push(all[i]); err != nil {
			t.Fatal(err)
		}
	}

	// Get all
	get, err := q.Get()
	if err != nil {
		t.Fatal(err)
	}
	if gLen, aLen := len(get), len(all); gLen != aLen {
		t.Fatalf("Length mismatch %d != %d", gLen, aLen)
	}
	for i := range all {
		if g, a := get[i], all[i]; g != a {
			t.Errorf("ID Mismatch %s != %s", g, a)
		}
	}
}

func TestPushUnshift(t *testing.T) {
	q := connect(t, "TestPushUnshift")
	defer close(q)
	// Reset/init ids
	if err := q.Put([]bson.ObjectId{}); err != nil {
		t.Fatal(err)
	}
	all := make([]bson.ObjectId, 10)
	for i := 0; i < 10; i++ {
		all[i] = bson.NewObjectId()
	}

	// Push
	for i := range all {
		if err := q.Push(all[i]); err != nil {
			t.Fatal(err)
		}
	}
	for i := range all {
		id, err := q.Unshift()
		if err != nil {
			t.Fatal(err)
		}
		if all[i] != id {
			t.Errorf("ID Mismatch %s != %s", all[i], id)
		}
	}
}

func TestRace(t *testing.T) {
	q := connect(t, "TestRace")
	defer close(q)
	// Reset/init ids
	all := make([]bson.ObjectId, 10)
	for i := range all {
		all[i] = bson.NewObjectId()
	}
	if err := q.Put(all); err != nil {
		t.Fatal(err)
	}

	// Get IDs all at once
	ch := make(chan bson.ObjectId, len(all))
	for _ = range all {
		go func() {
			id, err := q.Unshift()
			if err != nil {
				t.Error(err)
			}
			ch <- id
		}()
	}

	for _ = range all {
		select {
		case id := <-ch:
			t.Log(id)
		case <-time.After(10 * time.Second):
			t.Fatal("Timeout reached")
		}
	}
}

func connect(t *testing.T, name string) (q *IdQueue) {
	q = &IdQueue{
		Name: name,
		Addr: "localhost:8046",
	}
	if err := q.Connect(); err != nil {
		t.Fatal(err)
	}
	return
}

func close(q *IdQueue) {
	q.Close()
}
