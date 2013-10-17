package main

import (
	"flag"
	"fmt"
	"github.com/300brand/coverage/cleanurl"
	"net/url"
	"os"
)

func main() {
	flag.Parse()
	rawurl := flag.Arg(0)

	parsed, err := url.Parse(rawurl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cleaned := cleanurl.Clean(parsed)
	fmt.Println(cleaned)
}
