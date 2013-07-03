package search

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
)

type IdFilter struct {
	IdMap  map[bson.ObjectId]byte
	Target byte
}

func NewIdFilter(target int) (f *IdFilter) {
	return &IdFilter{
		IdMap:  make(map[bson.ObjectId]byte, 256),
		Target: byte(target),
	}
}

func (f *IdFilter) Add(kw *coverage.Keyword) {
	f.IdMap[kw.ArticleId]++
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
