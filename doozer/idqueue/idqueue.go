package idqueue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ha/doozer"
	"labix.org/v2/mgo/bson"
	"sync"
)

type IdQueue struct {
	Name   string
	Addr   string
	Max    int
	conn   *doozer.Conn
	file   string
	dec    *json.Decoder
	enc    *json.Encoder
	buf    *bytes.Buffer
	lowtex *sync.Mutex // Low-Level Mutex - Put/Get
	hitex  *sync.Mutex // High-Level Mutex - Push/Unshift
}

var (
	ErrEOQ  = errors.New("End of queue")
	ErrFull = errors.New("Queue full")
)

func (q *IdQueue) Connect() (err error) {
	if q.conn, err = doozer.Dial(q.Addr); err != nil {
		return
	}
	rev, err := q.conn.Rev()
	if q.Name == "" {
		q.Name = fmt.Sprintf("ids-%d", rev)
	}
	if q.Max == 0 {
		q.Max = 100
	}
	q.file = fmt.Sprintf("/queue/%s", q.Name)

	bufSize := 27*q.Max + 1 // Size of q.Max JSON-encoded BSON ObjectIds
	q.buf = bytes.NewBuffer(make([]byte, bufSize))
	q.dec, q.enc = json.NewDecoder(q.buf), json.NewEncoder(q.buf)

	q.lowtex = &sync.Mutex{}
	q.hitex = &sync.Mutex{}
	return
}

func (q *IdQueue) Close() {
	q.conn.Close()
}

func (q *IdQueue) Get() (ids []bson.ObjectId, err error) {
	q.lowtex.Lock()
	defer q.lowtex.Unlock()

	b, _, err := q.conn.Get(q.file, nil)
	if err != nil {
		return
	}
	q.buf.Reset()
	if _, err = q.buf.Write(b); err != nil {
		return
	}
	if err = q.dec.Decode(&ids); err != nil {
		return
	}
	return
}

func (q *IdQueue) Put(ids []bson.ObjectId) (err error) {
	q.lowtex.Lock()
	defer q.lowtex.Unlock()

	rev, err := q.conn.Rev()
	if err != nil {
		return err
	}
	q.buf.Reset()
	q.enc.Encode(ids)
	q.conn.Set(q.file, rev, q.buf.Bytes())
	return
}

func (q *IdQueue) Push(id bson.ObjectId) (err error) {
	q.hitex.Lock()
	defer q.hitex.Unlock()

	ids, err := q.Get()
	if err != nil {
		return
	}
	if len(ids) >= q.Max {
		return ErrFull
	}
	ids = append(ids, id)
	return q.Put(ids)
}

func (q *IdQueue) Unshift() (id bson.ObjectId, err error) {
	q.hitex.Lock()
	defer q.hitex.Unlock()

	ids, err := q.Get()
	if err != nil {
		return
	}
	if len(ids) == 0 {
		err = ErrEOQ
		return
	}
	id = ids[0]
	return id, q.Put(ids[1:])
}
