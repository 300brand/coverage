package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/parser"
	"git.300brand.com/coverage/parser/normalizer"
	"io/ioutil"
	"os"
	"path"
)

var local bool

func init() {
	flag.BoolVar(&local, "local", false, "Argument is a local file")
}

func DownloadFeed(url string) (data []byte, err error) {
	resp, err := downloader.Fetch(url)
	if err != nil {
		return
	}
	data = resp.Body
	return
}

func GetData(path string, local bool) (data []byte, err error) {
	if local {
		data, err = ReadFeed(path)
	} else {
		data, err = DownloadFeed(path)
	}
	return
}

func ReadFeed(path string) (data []byte, err error) {
	return ioutil.ReadFile(path)
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Printf("Usage: %s [-url] FEED\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	path := flag.Arg(0)

	// Pull data from download or filesystem
	data, err := GetData(path, local)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Normalize Feed
	n := &normalizer.Default{}
	if err := parser.Normalize(data, n); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Output as JSON
	enc, err := json.MarshalIndent(n, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	os.Stdout.Write(enc)
}
