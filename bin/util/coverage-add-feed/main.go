package main

import (
	"flag"
	"fmt"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"net/url"
	"os"
)

var (
	err     error
	feedURL string
	pubID   string
	f       *coverage.Feed
	p       *coverage.Publication
)

func main() {
	flag.Parse()

	feedURL = flag.Arg(0)
	pubID = flag.Arg(1)

	m := mongo.New("localhost", "Coverage")
	if err = m.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer m.Close()

	// Build feed
	f = coverage.NewFeed()
	if f.URL, err = url.Parse(feedURL); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if err = m.UpdateFeed(f); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Associate feed with publication
	if p, err = m.GetPublication(bson.M{"_id": bson.ObjectIdHex(pubID)}); err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	p.AddFeed(f)
	if err = m.UpdatePublication(p); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
}
