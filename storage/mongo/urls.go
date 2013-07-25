package mongo

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/url"
	"sync"
)

const URLsCollection = "URLs"

var (
	ErrURLExists = errors.New("URL exists")
	mutex        sync.Mutex
)

func (m *Mongo) AddURL(u *url.URL, id bson.ObjectId) (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Check for URL
	n, err := m.db.C(URLsCollection).Find(bson.M{"_id": u}).Limit(1).Count()
	if err != nil && err != mgo.ErrNotFound {
		return
	}
	if n == 1 {
		return ErrURLExists
	}

	// Insert!
	return m.db.C(URLsCollection).Insert(bson.M{"_id": u, "article": id})
}
