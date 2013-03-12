package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"net/url"
	"os"
)

var (
	err    error
	title  string
	pubURL string
	p      *coverage.Publication
)

func main() {
	flag.Parse()

	pubURL = flag.Arg(0)
	title = flag.Arg(1)

	m := mongo.New("localhost", "Coverage")
	if err = m.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer m.Close()

	p = coverage.NewPublication()

	p.Title = title

	if p.URL, err = url.Parse(pubURL); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	if err = m.UpdatePublication(p); err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	fmt.Println(p.ID.Hex())
}
