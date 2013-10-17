package main

import (
	"flag"
	"fmt"
	"github.com/300brand/coverage/article/lexer"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var config = struct {
	File bool
}{}

func init() {
	flag.BoolVar(&config.File, "f", false, "Read from file instead of stdin")
}

func main() {
	flag.Parse()

	var r io.Reader
	if config.File {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		r = f
	} else {
		r = os.Stdin
	}

	in, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range lexer.Words(in) {
		fmt.Printf("All: %s\n", w)
	}
	for _, w := range lexer.Keywords(in) {
		fmt.Printf("Keyword: %s\n", w)
	}
}
