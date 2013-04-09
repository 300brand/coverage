package mongo

import (
	"git.300brand.com/coverage/bridge"
	"time"
)

type queueState struct {
	LastId  uint64
	LastRun time.Time
}

const QueueStateCollection = "QueueState"

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
	state := &queueState{q.LastId, time.Now()}
	_, err = m.db.C(QueueStateCollection).Upsert(nil, state)
	return
}
