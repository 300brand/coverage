package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"hash/fnv"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
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
	hasher  = fnv.New32a()
	layout  = "2006.01.02-15.04.05"
	from    = time.Now().Add(-1 * time.Hour)
	to      = time.Now()
	dbHost  = flag.String("dbHost", "localhost", "MongoDB host")
	dbName  = flag.String("dbName", "Coverage", "MongoDB database")
	strFrom = flag.String("from", from.Format(layout), "From search bounds")
	strTo   = flag.String("to", to.Format(layout), "To search bounds")
	toJSON  = flag.Bool("json", false, "Print article IDs as a JSON array")
	stats   = Stats{
		BatchSize:  0,
		BatchStart: time.Now(),
		Start:      time.Now(),
		Frequency:  15 * time.Second,
	}
)

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
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

	KC := mongos.DB("Keywords").C("Keywords2")
	// if err := KC.Create(&mgo.CollectionInfo{DisableIdIndex: true}); err != nil {
	// 	log.Print(err)
	// }
	// KC.EnsureIndex(mgo.Index{
	// 	Key:    []string{"_id", "date"},
	// 	Unique: true,
	// })

	AC := session.DB(*dbName).C("Articles")
	query := bson.M{
		"_id": bson.M{
			"$gte": bson.NewObjectIdWithTime(from),
			"$lte": bson.NewObjectIdWithTime(to),
		},
	}
	sel := bson.M{
		"_id": 1,
		"text.words.keywords": 1,
		"published":           1,
		"added":               1,
	}
	log.Printf("Query: %v", query)
	iter := AC.Find(query).Select(sel).Iter()

	var d time.Time
	threshold := time.Unix(0, 0)
	aCh := make(chan coverage.Article, BufferSize)
	readyChan := make(chan bool)

	go func() {
		a := &coverage.Article{}
		count := 0
		fmt.Print("Initial Buffer ")
		for iter.Next(a) {
			aCh <- *a
			count++
			fmt.Print(".")

			if count%BufferSize == 0 {
				fmt.Println("")
				readyChan <- true
				fmt.Print("Buffering ")
			}
		}
		close(aCh)
		if err = iter.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for a := range aCh {
		stats.ArticleCount++
		if a.Published.After(threshold) {
			d = a.Published.Truncate(24 * time.Hour)
		} else {
			d = a.Added.Truncate(24 * time.Hour)
		}

		for _, w := range a.Text.Words.Keywords {
			stats.KeywordCount++

			hash := coverage.KeywordHash(w)

			find := bson.M{"hash": hash, "date": d}
			change := bson.M{
				"$setOnInsert": bson.M{
					"hash":    hash,
					"keyword": w,
					"date":    d,
				},
				"$push": bson.M{
					"articles": a.ID,
				},
			}
			if _, err := KC.Upsert(find, change); err != nil {
				log.Fatal(err)
			}
			// switch err := KC.Update(find, change); err {
			// case mgo.ErrNotFound:
			// 	kw := &coverage.Keyword{
			// 		Hash:     hash,
			// 		Keyword:  w,
			// 		Date:     d,
			// 		Articles: []bson.ObjectId{a.ID},
			// 	}
			// 	if err := KC.Insert(kw); err != nil {
			// 		log.Fatal(err)
			// 	}
			// case nil:
			// default:
			// 	log.Fatal(err)
			// }
		}

		if stats.ArticleCount%BufferSize == 0 {
			<-readyChan
		}
		stats.Print()
	}

	fmt.Printf("Imported %d Articles and %d Keywords\n", stats.ArticleCount, stats.KeywordCount)
	fmt.Printf("Completed in %s\n", time.Since(stats.Start))
}
