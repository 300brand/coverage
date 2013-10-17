package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/300brand/coverage"
	"hash/fnv"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"sync"
	"time"
)

type Stats struct {
	ArticleCount int
	KeywordCount int
	BatchSize    int
	Start        time.Time
	BatchStart   time.Time
	Frequency    time.Duration
}

const BufferSize = 500

var (
	hasher    = fnv.New32a()
	layout    = "2006-01-02"
	from      = time.Now().Add(-1 * time.Hour)
	to        = time.Now()
	dbHost    = flag.String("dbHost", "localhost", "MongoDB host")
	dbName    = flag.String("dbName", "Coverage", "MongoDB database")
	strFrom   = flag.String("from", from.Format(layout), "From search bounds")
	strTo     = flag.String("to", to.Format(layout), "To search bounds")
	toJSON    = flag.Bool("json", false, "Print article IDs as a JSON array")
	threshold = time.Unix(0, 0)
	stats     = Stats{
		BatchSize:  0,
		BatchStart: time.Now(),
		Start:      time.Now(),
		Frequency:  15 * time.Second,
	}
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func (s *Stats) Print() {
	s.BatchSize++
	if time.Since(s.BatchStart) < s.Frequency {
		return
	}
	fmt.Printf("Articles: %6d Keywords: %8d Rate: %d/%s\n", s.ArticleCount, s.KeywordCount, s.BatchSize, time.Since(s.BatchStart))
	s.BatchStart = time.Now()
	s.BatchSize = 0
}

func main() {
	var err error
	flag.Parse()

	if from, err = time.ParseInLocation(layout, *strFrom, time.Local); err != nil {
		log.Fatal(err)
	}

	if to, err = time.ParseInLocation(layout, *strTo, time.Local); err != nil {
		log.Fatal(err)
	}

	if from.After(to) {
		log.Fatal("from date comes before to date")
	}

	session, err := mgo.Dial(*dbHost)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	mongos, err := mgo.Dial("data0.coverage.net:27020")
	if err != nil {
		log.Fatal(err)
	}
	defer mongos.Close()

	KC := mongos.DB("Keywords").C("Keywords")
	// if err := KC.Create(&mgo.CollectionInfo{DisableIdIndex: true}); err != nil {
	// 	log.Print(err)
	// }
	// KC.EnsureIndex(mgo.Index{
	// 	Key:    []string{"_id", "date"},
	// 	Unique: true,
	// })

	AC := session.DB(*dbName).C("Articles")
	sel := bson.M{
		"_id": 1,
		"text.words.keywords": 1,
		"published":           1,
		"added":               1,
	}
	a := &coverage.Article{}
	//hashes := make(map[string]uint32, 4096)
	var wg sync.WaitGroup

	for t := to; t.After(from); t = t.AddDate(0, 0, -1) {
		date := t.AddDate(0, 0, -1)
		query := bson.M{
			"published": bson.M{
				"$gte": date,
				"$lt":  t,
			},
		}
		log.Printf("Date: %s Query:", date)
		json.NewEncoder(os.Stdout).Encode(query)

		kws := make(map[string][]bson.ObjectId, 4096)

		iter := AC.Find(query).Select(sel).Iter()
		for iter.Next(a) {
			stats.ArticleCount++
			for _, w := range a.Text.Words.Keywords {
				stats.KeywordCount++
				//if _, ok := hashes[w]; !ok {
				//	hashes[w] = coverage.KeywordHash(w)
				//}
				kws[w] = append(kws[w], a.ID)
			}
			stats.Print()
		}

		wg.Add(1)
		go func(kws map[string][]bson.ObjectId, date time.Time) {
			log.Printf("[%s] Inserting Keywords...", date)
			start := time.Now()
			docs := make([]interface{}, 0, len(kws))
			for w, ids := range kws {
				docs = append(docs, coverage.Keyword{
					Id: coverage.KeywordId{
						//Hash: hashes[w],
						Date:    date,
						Keyword: w,
					},
					Articles: ids,
				})
			}
			for i, l := 0, 1000; i < len(docs); i, l = l, l+1000 {
				if l > len(docs) {
					l = len(docs)
				}
				if err := KC.Insert(docs[i:l]...); err != nil {
					//fmt.Printf("%+v\n", hashes)
					log.Fatal(err)
				}
				log.Printf("[%s] Inserted %d Keyword docs. Took %s", date, len(docs[i:l]), time.Since(start))
			}
			wg.Done()
		}(kws, date)

	}

	log.Println("Waiting for inserts to complete...")
	wg.Wait()

	fmt.Printf("Imported %d Articles and %d Keywords\n", stats.ArticleCount, stats.KeywordCount)
	fmt.Printf("Completed in %s\n", time.Since(stats.Start))
}
