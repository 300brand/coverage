package main

/*
 * Redacted for later refinement
import (
	"flag"
	"git.300brand.com/coverage/bridge"
	"git.300brand.com/coverage/storage/mongo"
	"log"
)

var limit uint64

func init() {
	flag.Uint64Var(&limit, "n", 10, "Chunk size for pulling items from frontend queue")
}
*/
func main() {
	/*
		flag.Parse()

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

		queue, err := bridge.GetQueue(lastId, limit)
		if err != nil {
			log.Fatalf("bridge.GetQueue: %s", err)
		}

		fs := mongo.NewFeedService(m)
		for _, f := range queue.NewFeeds {
			log.Printf("Adding Feed: %s %s", f.ID.Hex(), f.URL)
			if err := fs.Update(&f); err != nil {
				log.Printf("fs.Update: %s", err)
			}
		}

		rs := mongo.NewReportService(m)
		for _, r := range queue.Reports {
			log.Printf("Adding Report: %s", r.ID.Hex())
			if err := rs.Update(&r); err != nil {
				log.Printf("rs.Update: %s", err)
			}
		}

		if err := q.UpdateQueue(queue); err != nil {
			log.Fatalf("q.UpdateQueue: %s", err)
		}
	*/
}
