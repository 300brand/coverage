package main

import (
	"encoding/json"
	"flag"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"io"
	"log"
	"net/url"
	"os"
)

type Pubs []Pub

type Pub struct {
	Title string
	URL   string
	Feeds []string
}

var (
	dbHost = "localhost"
	dbName = "Coverage"
)

func init() {
	flag.StringVar(&dbHost, "dbHost", dbHost, "Override MongoDB host")
	flag.StringVar(&dbName, "dbName", dbName, "Override MongoDB database")
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
		}
		for _, feedURL := range pub.Feeds {
			f := coverage.NewFeed()
			if f.URL, err = url.Parse(feedURL); err != nil {
				p.AddFeed(f)
			}
			if err = feedService.Update(f); err != nil {
				log.Println(err)
			}
		}
		if err = pubService.Update(p); err != nil {
			log.Println(err)
		}
	}
}
