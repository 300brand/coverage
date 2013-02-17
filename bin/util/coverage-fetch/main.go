package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage/downloader"
	"os"
)

func main() {
	flag.Parse()

	url := flag.Arg(0)
	r, err := downloader.Fetch(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Stdout.Write(r.Body)
}
