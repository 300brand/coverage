package main

import (
	"flag"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/article/lexer"
	"git.300brand.com/coverage/search"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	layout = "2006-01-02"
	from   time.Time
	to     time.Time
	wg     sync.WaitGroup
	config = struct {
		ArticleHost *string
		ArticleDB   *string
		KeywordHost *string
		KeywordDB   *string
		KeywordColl *string
		From        *string
		To          *string
	}{
		flag.String("articles.host", "localhost", "Articles MongoDB host"),
		flag.String("articles.db", "Coverage", "Articles database"),
		flag.String("keywords.host", "mongos0.coverage.net:27020", "Keywords MongoDB host"),
		flag.String("keywords.db", "Keywords", "Keywords database"),
		flag.String("keywords.coll", "Keywords3", "Keywords collection"),
		flag.String("from", "2013-01-01", "From search bounds"),
		flag.String("to", "2013-04-09", "To search bounds"),
	}
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	var err error

	flag.Parse()

	loc := time.Local

	if from, err = time.ParseInLocation(layout, *config.From, loc); err != nil {
		log.Fatal(err)
	}
	if to, err = time.ParseInLocation(layout, *config.To, loc); err != nil {
		log.Fatal(err)
	}

	articleDb := mongo.New(*config.ArticleHost, *config.ArticleDB)
	if err := articleDb.Connect(); err != nil {
		log.Fatal(err)
	}
	defer articleDb.Close()

	kSess, err := mgo.Dial(*config.KeywordHost)
	if err != nil {
		log.Fatal(err)
	}
	defer kSess.Close()
	kColl := kSess.DB(*config.KeywordDB).C(*config.KeywordColl)

	q := strings.Join(flag.Args(), " ")
	log.Printf("Search Query: '%s'", q)
	terms := lexer.Keywords([]byte(q))
	idChan := make(chan bson.ObjectId, 100)
	ids := search.NewIdFilter(len(terms))
	ids.UseChan()
	boolChan := make(chan *coverage.Article)
	boolSearch := search.NewBoolean(q)

	go func() {
		for id := range idChan {
			ids.Add(id)
		}
		close(ids.Chan)
	}()

	go func() {
		all := search.All{}
		var wait sync.WaitGroup
		for id := range ids.Chan {
			wait.Add(1)
			go func(id bson.ObjectId) {
				defer wait.Done()
				a, err := articleDb.GetArticle(bson.M{"_id": id})
				if err != nil {
					log.Fatal(err)
				}
				stats := &search.Stats{}
				if err := all.Matches(a, terms, stats); err != nil {
					log.Print(err)
				}
				//log.Printf("%s All: %v", a.ID, stats.All)
				if !stats.All {
					return
				}
				if !boolSearch.Match(a.Text.Body.Text) {
					return
				}
				boolChan <- a
			}(id)
		}
		wait.Wait()
		close(boolChan)
	}()

	for _, term := range terms {
		for t := to; t.After(from.AddDate(0, 0, -1)); t = t.AddDate(0, 0, -1) {
			wg.Add(1)
			go func(term string, t time.Time) {
				log.Printf("    -> %s %s", t.Format(layout), term)
				kw := &coverage.Keyword{}
				find := bson.M{"_id": bson.M{"keyword": term, "date": t}}
				if err := kColl.Find(find).One(kw); err != nil {
					log.Print(err)
				}
				for _, id := range kw.Articles {
					idChan <- id
				}
				log.Printf("    <- %s %s", t.Format(layout), term)
				wg.Done()
			}(term, t)
		}
	}

	log.Println("Waiting for idChan to finish")
	wg.Wait()
	close(idChan)

	for a := range boolChan {
		log.Printf("BOOL MATCH: %s", a.ID)
	}

	//log.Println(ids.Ids())

	// if remap {
	// 	start := bson.NewObjectIdWithTime(from)
	// 	end := bson.NewObjectIdWithTime(to)

	// 	log.Printf("Starting keyword map-reduce between %s and %s", start, end)
	// 	bounds := bson.M{
	// 		"_id": bson.M{
	// 			"$lte": end,
	// 			"$gte": start,
	// 		},
	// 	}
	// 	info, err := m.ReduceKeywords(bounds)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("InputCount:  %d", info.InputCount)
	// 	log.Printf("EmitCount:   %d", info.EmitCount)
	// 	log.Printf("OutputCount: %d", info.OutputCount)
	// 	log.Printf("EmitLoop:    %s", time.Duration(info.VerboseTime.EmitLoop))
	// 	log.Printf("Map:         %s", time.Duration(info.VerboseTime.Map))
	// 	log.Printf("Time:        %s", time.Duration(info.Time))
	// }

	// now := time.Now()
	// terms := flag.Args()
	// count := 0
	// kwChan := make(chan coverage.Keyword)
	// filter := search.NewIdFilter(len(terms))
	// go m.KeywordSearch(terms, from, to, kwChan)

	// for kw := range kwChan {
	// 	filter.Add(&kw)
	// 	count++
	// }
	// log.Printf("Found %d Article IDs matching ANY %v in %s", count, terms, time.Since(now))

	// now = time.Now()
	// ids := filter.Ids()
	// log.Printf("Found %d Article IDs matching ALL %v in %s", len(ids), terms, time.Since(now))

	// for i := 0; i < 100 && i < len(ids); i++ {
	// 	a, err := m.GetArticle(bson.M{"_id": ids[i]})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(a.URL)
	// }

	// if *toJSON {
	// 	enc := json.NewEncoder(os.Stdout)
	// 	if err := enc.Encode(ids); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
