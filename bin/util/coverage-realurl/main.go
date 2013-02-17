package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage/downloader"
	"os"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No URL supplied")
		os.Exit(1)
	}

	url := flag.Arg(0)
	r, err := downloader.Fetch(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(r.RealURL)
}
