package mongo

import (
	"labix.org/v2/mgo/bson"
)

const URLsCollection = "URLs"

func (m *Mongo) AddURL(u string, id bson.ObjectId) (err error) {
	c := m.Copy()
	defer c.Close()

	return c.URLs.Insert(bson.M{"_id": u, "article": id})
}
