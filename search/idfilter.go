package search

import (
	"labix.org/v2/mgo/bson"
)

type IdFilter struct {
	IdMap  map[bson.ObjectId]byte
	Target byte
	Chan   chan bson.ObjectId
}

func NewIdFilter(target int) (f *IdFilter) {
	return &IdFilter{
		IdMap:  make(map[bson.ObjectId]byte, 256),
		Target: byte(target),
	}
}

func (f *IdFilter) Add(id bson.ObjectId) {
	f.IdMap[id]++
	if f.Chan != nil && f.IdMap[id] == f.Target {
		f.Chan <- id
	}
}

func (f *IdFilter) Ids() (ids []bson.ObjectId) {
	ids = make([]bson.ObjectId, 0, len(f.IdMap))
	for id, c := range f.IdMap {
		if c == f.Target {
			ids = append(ids, id)
		}
	}
	return
}

func (f *IdFilter) UseChan() {
	f.Chan = make(chan bson.ObjectId)
}
