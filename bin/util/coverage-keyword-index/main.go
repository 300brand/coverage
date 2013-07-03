package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type Keyword struct {
	Hash      uint32 `bson:",minsize"`
	Keyword   string
	Date      time.Time
	ArticleId bson.ObjectId
	Published time.Time `bson:"-"`
}

type Stats struct {
	ArticleCount int
	KeywordCount int
	BatchSize    int
	Start        time.Time
	BatchStart   time.Time
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

var (
	layout  = "2006.01.02-15.04.05"
	from    = time.Now().Add(-1 * time.Hour)
	to      = time.Now()
	dbHost  = flag.String("dbHost", "localhost", "MongoDB host")
	dbName  = flag.String("dbName", "Coverage", "MongoDB database")
	strFrom = flag.String("from", from.Format(layout), "From search bounds")
	strTo   = flag.String("to", to.Format(layout), "To search bounds")
	toJSON  = flag.Bool("json", false, "Print article IDs as a JSON array")
	stats   = Stats{
		BatchSize:  1000,
		BatchStart: time.Now(),
		Start:      time.Now(),
	}
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func (s *Stats) Print() {
	if s.ArticleCount%s.BatchSize != 0 {
		return
	}
	fmt.Printf("Articles: %6d Keywords: %8d Rate: %d/%s\n", s.ArticleCount, s.KeywordCount, s.BatchSize, time.Since(s.BatchStart))
	s.BatchStart = time.Now()
}

func hashRK(s string) (hash uint32) {
	for i := 0; i < len(s); i++ {
		hash = hash*primeRK + uint32(s[i])
	}
	return
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

	KC := session.DB(*dbName + "_keywords").C("Keywords")
	err = KC.Create(&mgo.CollectionInfo{DisableIdIndex: true})
	if err != nil {
		log.Print(err)
	}
	KC.EnsureIndexKey("hash")
	//KC.EnsureIndexKey("date")

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

	kw := &Keyword{}
	kws := make([]interface{}, 512)
	a := &coverage.Article{}
	threshold := time.Unix(0, 0)

	for iter.Next(a) {
		stats.ArticleCount++
		kw.ArticleId = a.ID
		if a.Published.After(threshold) {
			kw.Published = a.Published
		} else {
			kw.Published = a.Added
		}
		kw.Date = kw.Published.Truncate(24 * time.Hour)
		if l := len(a.Text.Words.Keywords); cap(kws) < l {
			kws = make([]interface{}, l)
		}
		for i, w := range a.Text.Words.Keywords {
			stats.KeywordCount++
			kw.Hash = hashRK(w)
			kw.Keyword = w
			kws[i] = *kw
		}
		if err = KC.Insert(kws[:len(a.Text.Words.Keywords)]...); err != nil {
			log.Print(err)
		}
		stats.Print()
	}

	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Imported %d Articles and %d Keywords\n", stats.ArticleCount, stats.KeywordCount)
	fmt.Printf("Completed in %s\n", time.Since(stats.Start))
}
