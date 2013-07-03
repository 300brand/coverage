package main

import (
	"math"
	"encoding/json"
	"flag"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/search"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"time"
)

var (
	layout                     = "2006.01.02-15.04.05"
	dbHost, dbName, dbKeywords string
	remap                      bool
	from                       = time.Now().Add(-1 * time.Hour)
	to                         = time.Now()
	strFrom                    = flag.String("from", from.Format(layout), "From search bounds")
	strTo                      = flag.String("to", to.Format(layout), "To search bounds")
	toJSON                     = flag.Bool("json", false, "Print article IDs as a JSON array")
)

func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "MongoDB host")
	flag.StringVar(&dbName, "dbName", "Coverage", "MongoDB database")
	flag.StringVar(&dbKeywords, "dbKeywords", "Coverage", "MongoDB database")
	flag.BoolVar(&remap, "remap", false, "Re-run Map-Reduce")
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	var err error

	flag.Parse()

	loc := time.Local

	if from, err = time.ParseInLocation(layout, *strFrom, loc); err != nil {
		log.Fatal(err)
	}
	if to, err = time.ParseInLocation(layout, *strTo, loc); err != nil {
		log.Fatal(err)
	}

	m := mongo.New(dbHost, dbName)
	m.KeywordDB = "Coverage_keywords"
	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if remap {
		start := bson.NewObjectIdWithTime(from)
		end := bson.NewObjectIdWithTime(to)

		log.Printf("Starting keyword map-reduce between %s and %s", start, end)
		bounds := bson.M{
			"_id": bson.M{
				"$lte": end,
				"$gte": start,
			},
		}
		info, err := m.ReduceKeywords(bounds)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("InputCount:  %d", info.InputCount)
		log.Printf("EmitCount:   %d", info.EmitCount)
		log.Printf("OutputCount: %d", info.OutputCount)
		log.Printf("EmitLoop:    %s", time.Duration(info.VerboseTime.EmitLoop))
		log.Printf("Map:         %s", time.Duration(info.VerboseTime.Map))
		log.Printf("Time:        %s", time.Duration(info.Time))
	}

	now := time.Now()
	terms := flag.Args()
	count := 0
	kwChan := make(chan coverage.Keyword)
	filter := search.NewIdFilter(len(terms))
	go m.KeywordSearch(terms, from, to, kwChan)

	for kw := range kwChan {
		filter.Add(&kw)
		count++
	}
	log.Printf("Found %d Article IDs matching ANY %v in %s", count, terms, time.Since(now))

	now = time.Now()
	ids := filter.Ids()
	log.Printf("Found %d Article IDs matching ALL %v in %s", len(ids), terms, time.Since(now))

	for i := 0; i < 100 && i < len(ids); i++ {
		a, err := m.GetArticle(bson.M{"_id": ids[i]})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(a.URL)
	}

	if *toJSON {
		enc := json.NewEncoder(os.Stdout)
		if err := enc.Encode(ids); err != nil {
			log.Fatal(err)
		}
	}
}
