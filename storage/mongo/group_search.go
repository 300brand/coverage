package mongo

import (
	"labix.org/v2/mgo/bson"
	"github.com/300brand/coverage"
)

const GroupSearchCollection = "GroupSearch"

func (m *Mongo) GetGroupSearch(id bson.ObjectId, s *coverage.GroupSearch) (err error) {
	c := m.Copy()
	defer c.Close()
	return c.GroupSearch.FindId(id).One(s)
}

func (m *Mongo) UpdateGroupSearch(s *coverage.GroupSearch) (err error) {
	c := m.Copy()
	defer c.Close()
	_, err = c.GroupSearch.UpsertId(s.Id, s)
	return
}
