package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage/article/lexer"
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

	//lexer.StemmingEnabled = true
	words := lexer.GetWords(in)
	for _, w := range words {
		fmt.Printf("%-3d %-15s %-15s\n", w.Index, w.Word, w.Stem)
	}
}
