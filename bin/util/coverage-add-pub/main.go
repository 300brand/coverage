package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"net/url"
	"os"
)

var (
	err    error
	id     string
	title  string
	pubURL string
	p      *coverage.Publication
)

func init() {
	flag.StringVar(&id, "id", id, "Publication ID - set this to modify an existing publication")
	flag.StringVar(&title, "title", title, "Publication title")
	flag.StringVar(&pubURL, "url", pubURL, "Publication URL")
}

func main() {
	flag.Parse()

	if id == "" && title == "" {
		fmt.Println("Title required if ID not supplied")
		os.Exit(1)
	}
	if id == "" && pubURL == "" {
		fmt.Println("URL required if ID not supplied")
		os.Exit(2)
	}

	m := mongo.New("localhost", "Coverage")
	if err = m.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer m.Close()

	if id == "" {
		p = coverage.NewPublication()
	} else {
		if p, err = m.GetPublication(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	}

	if title != "" {
		p.Title = title
	}

	if p.URL, err = url.Parse(pubURL); pubURL != "" && err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	if err = m.UpdatePublication(p); err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	fmt.Println(p.ID.Hex())
}
