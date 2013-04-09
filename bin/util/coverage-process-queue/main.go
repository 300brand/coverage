package main

import (
	"git.300brand.com/coverage/bridge"
	"git.300brand.com/coverage/storage/mongo"
	"log"
)

func main() {
	m := mongo.New("localhost", "CoverageQueue")
	if err := m.Connect(); err != nil {
		log.Fatalf("m.Connect: %s", err)
	}
	defer m.Close()

	lastId, err := m.QueueLastId()
	if err != nil {
		log.Fatalf("m.QueueLastId: %s", err)
	}

	queue, err := bridge.GetQueue(lastId, 1)
	if err != nil {
		log.Fatalf("bridge.GetQueue: %s", err)
	}

	log.Printf("%+v", queue)

	if err := m.UpdateQueue(queue); err != nil {
		log.Fatalf("m.UpdateQueue: %s", err)
	}
}
