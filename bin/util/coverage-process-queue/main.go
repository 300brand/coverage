package main

import (
	"git.300brand.com/coverage/bridge"
	"git.300brand.com/coverage/storage/mongo"
	"log"
)

func main() {
	m := mongo.New("localhost", "CoverageQueue")
	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	lastId, err := m.QueueLastId()
	if err != nil {
		log.Fatal(err)
	}

	queue, err := bridge.GetQueue(lastId, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", queue)

	m.UpdateQueue(queue)
}
