package main

import (
	"encoding/json"
	"flag"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"io"
	"labix.org/v2/mgo/bson"
	"log"
	"net/url"
	"os"
)

type Pubs []Pub

type Pub struct {
	Title string
	URL   string
	Feeds []Feed
}

type Feed struct {
	ID  int
	URL string
}

type FeedID struct {
	MysqlID int
	MongoID bson.ObjectId
}

var (
	dbHost  string
	dbName  string
	feedIDs = []FeedID{}
)

func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "Override MongoDB host")
	flag.StringVar(&dbName, "dbName", "Coverage", "Override MongoDB database")
}

func main() {
	flag.Parse()

	var r io.Reader
	if flag.NArg() == 0 {
		r = os.Stdin
	} else {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	}

	pubs := &Pubs{}
	d := json.NewDecoder(r)
	if err := d.Decode(pubs); err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println("Setting up mongo services")
	m := mongo.New(dbHost, dbName)
	if err := m.Connect(); err != nil {
		log.Println(err)
		os.Exit(3)
	}
	defer m.Close()
	pubService := mongo.NewPublicationService(m)
	feedService := mongo.NewFeedService(m)

	var err error
	for _, pub := range *pubs {
		log.Printf("Adding Publication: %s [%s]", pub.Title, pub.URL)
		p := coverage.NewPublication()
		p.Title = pub.Title
		if p.URL, err = url.Parse(pub.URL); err != nil {
			log.Println(err)
			continue
		}
		for _, feed := range pub.Feeds {
			f := coverage.NewFeed()
			if f.URL, err = url.Parse(feed.URL); err != nil {
				log.Println(err)
				continue
			}
			p.AddFeed(f)
			if err = feedService.Update(f); err != nil {
				log.Println(err)
			}
			feedIDs = append(feedIDs, FeedID{
				MysqlID: feed.ID,
				MongoID: f.ID,
			})
		}
		if err = pubService.Update(p); err != nil {
			log.Println(err)
		}
	}

	b, err := json.MarshalIndent(&feedIDs, "", "\t")
	if err != nil {
		log.Println(err)
	}
	os.Stdout.Write(b)
}
