package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/300brand/coverage/article/body"
	"io/ioutil"
	"os"
)

var (
	cleanedOnly bool
	showHTML    bool
)

func init() {
	flag.BoolVar(&cleanedOnly, "clean", false, "Only show the cleaned HTML used before extracing body content")
	flag.BoolVar(&showHTML, "html", false, "Show HTML body instead of Text body")
}

func main() {
	flag.Parse()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	clean, err := body.CleanHTML(in)
	if cleanedOnly {
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Printf("%s\n", clean)
		return
	}

	b, err := body.GetBody(clean)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	if showHTML {
		fmt.Printf("%s\n", b.HTML)
	} else {
		fmt.Printf("%s\n", doubleSpace(b.Text))
	}
}

func doubleSpace(b []byte) []byte {
	return bytes.Replace(b, []byte("\n"), []byte("\n\n"), -1)
}
