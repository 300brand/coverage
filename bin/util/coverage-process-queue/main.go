package main

import (
	"git.300brand.com/coverage/bridge"
	"git.300brand.com/coverage/storage/mongo"
	"log"
)

func main() {
	q := mongo.New("localhost", "CoverageQueue")
	if err := q.Connect(); err != nil {
		log.Fatalf("q.Connect: %s", err)
	}
	defer q.Close()

	m := mongo.New("localhost", "Coverage")
	if err := m.Connect(); err != nil {
		log.Fatalf("m.Connect: %s", err)
	}
	defer m.Close()

	lastId, err := q.QueueLastId()
	if err != nil {
		log.Fatalf("q.QueueLastId: %s", err)
	}

	queue, err := bridge.GetQueue(lastId, 1)
	if err != nil {
		log.Fatalf("bridge.GetQueue: %s", err)
	}

	rs := mongo.NewReportService(m)
	for _, r := range queue.Reports {
		log.Printf("Adding Report: %s", r.ID.Hex())
		if err := rs.Update(&r); err != nil {
			log.Fatalf("rs.Update: %s", err)
		}
	}

	if err := q.UpdateQueue(queue); err != nil {
		log.Fatalf("q.UpdateQueue: %s", err)
	}
}
