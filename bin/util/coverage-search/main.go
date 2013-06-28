package main

import (
	"encoding/json"
	"flag"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"time"
)

var (
	layout         = "2006.01.02-15.04.05"
	dbHost, dbName string
	remap          bool
	from           = time.Now().Add(-1 * time.Hour)
	to             = time.Now()
	strFrom        = flag.String("from", from.Format(layout), "From search bounds")
	strTo          = flag.String("to", to.Format(layout), "To search bounds")
)

func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "MongoDB host")
	flag.StringVar(&dbName, "dbName", "Coverage", "MongoDB database")
	flag.BoolVar(&remap, "remap", false, "Re-run Map-Reduce")
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	var err error

	flag.Parse()

	if from, err = time.Parse(layout, *strFrom); err != nil {
		log.Fatal(err)
	}
	if to, err = time.Parse(layout, *strTo); err != nil {
		log.Fatal(err)
	}

	m := mongo.New(dbHost, dbName)
	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if remap {
		start := bson.NewObjectIdWithTime(from)
		end := bson.NewObjectIdWithTime(to)

		log.Printf("Starting keyword map-reduce between %s and %s", from, to)
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

	ids, err := m.KeywordArticleIds(flag.Args(), from, to)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(ids); err != nil {
		log.Fatal(err)
	}
}
