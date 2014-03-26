package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/downloader"
	"github.com/300brand/coverage/feed"
	"log"
	"net/url"
	"os"
)

var local bool

func init() {
	flag.BoolVar(&local, "local", false, "Argument is a local file")
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	path := flag.Arg(0)

	if local {
		cwd, _ := os.Getwd()
		path = fmt.Sprintf("file://%s/%s", cwd, path)
	}

	log.Printf("Downloading %s", path)

	u, err := url.Parse(path)
	if err != nil {
		log.Fatalf("url.Parse: %s", err)
	}

	f := coverage.NewFeed()
	f.URL = u.String()

	if err := downloader.NewFeedService().Update(f); err != nil {
		log.Fatalf("downloader: %s", err)
	}

	if err := feed.NewFeedService().Update(f); err != nil {
		log.Fatalf("feed: %s", err)
	}

	// Output as JSON
	enc, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	os.Stdout.Write(enc)
	os.Stdout.Write([]byte{'\n'})

	log.Print("Done")
}
