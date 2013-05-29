package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"sync"
)

// Mongo reports this size for CappedDoc documents
const CappedDocSize = 68

type Capped struct {
	Collection *mgo.Collection
	mutex      sync.Mutex
}

type CappedDoc struct {
	ObjectId   bson.ObjectId `bson:"_id,omitempty"`
	Id         bson.ObjectId
	Priority   int
	Processing bool
}

func (m *Mongo) CappedCollection(name string, maxDocs int) (c *Capped, err error) {
	c = &Capped{
		Collection: m.db.C(name),
	}

	names, err := m.db.CollectionNames()
	if err != nil {
		return
	}
	create := true
	for _, n := range names {
		if n == name {
			create = false
		}
	}

	if create {
		err = c.Collection.Create(&mgo.CollectionInfo{
			Capped:   true,
			MaxBytes: maxDocs * CappedDocSize,
			MaxDocs:  maxDocs,
		})
	}
	return
}

func (c *Capped) Insert(id bson.ObjectId, priority int) error {
	return c.Collection.Insert(&CappedDoc{Id: id, Priority: priority})
}

func (c *Capped) NextId() (id bson.ObjectId, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	doc := &CappedDoc{}
	if err = c.Collection.Find(bson.M{"processing": false}).One(doc); err != nil {
		return
	}
	if err = c.Collection.UpdateId(doc.ObjectId, bson.M{"$set": bson.M{"processing": true}}); err != nil {
		return
	}
	return doc.Id, nil
}
