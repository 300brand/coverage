package mongo

import (
	"labix.org/v2/mgo/bson"
	"net/url"
)

const URLsCollection = "URLs"

func (m *Mongo) AddURL(u *url.URL, id bson.ObjectId) (err error) {
	return m.C.URLs.Insert(bson.M{"_id": u, "article": id})
}
