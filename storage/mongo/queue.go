package mongo

import (
	//"git.300brand.com/coverage/bridge"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"time"
)

type queueState struct {
	LastId  uint64
	LastRun time.Time
}

const (
	QueueStateCollection = "QueueState"
	QueueLogCollection   = "QueueLog"
)

/*
func (m *Mongo) QueueLastId() (id uint64, err error) {
	s := &queueState{}
	q := m.db.C(QueueStateCollection).Find(nil)

	n, err := q.Count()
	if err != nil {
		return
	}

	// new/empty collection
	if n == 0 {
		return 0, nil
	}

	err = q.One(s)
	return s.LastId, err
}

func (m *Mongo) UpdateQueue(q bridge.Queue) (err error) {
	var c *mgo.Collection

	c = m.db.C(QueueLogCollection + ".NewFeeds")
	for _, f := range q.NewFeeds {
		if _, err = c.Upsert(bson.M{"queueid": f.QueueId}, f); err != nil {
			return
		}
	}

	c = m.db.C(QueueLogCollection + ".RemoveFeeds")
	for _, f := range q.RemoveFeeds {
		if _, err = c.Upsert(bson.M{"queueid": f.QueueId}, f); err != nil {
			return
		}
	}

	c = m.db.C(QueueLogCollection + ".Reports")
	for _, f := range q.Reports {
		if _, err = c.Upsert(bson.M{"queueid": f.QueueId}, f); err != nil {
			return
		}
	}

	state := &queueState{q.LastId, time.Now()}
	_, err = m.db.C(QueueStateCollection).Upsert(nil, state)
	return
}
*/
