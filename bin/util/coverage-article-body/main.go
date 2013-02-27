package main

import (
	"bytes"
	"flag"
	"fmt"
	"git.300brand.com/coverage/article/body"
	"io/ioutil"
	"os"
)

var (
	cleanedOnly bool
	noHTML      bool
	noText      bool
)

func init() {
	flag.BoolVar(&cleanedOnly, "clean", false, "Only show the cleaned HTML used before extracing body content")
	flag.BoolVar(&noHTML, "nohtml", false, "Don't show body HTML")
	flag.BoolVar(&noText, "notext", false, "Don't show body Text")
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
	if !noHTML {
		fmt.Printf("%s\n", b.HTML)
	}
	if !noHTML && !noText {
		fmt.Println("\n========================================\n")
	}
	if !noText {
		fmt.Printf("%s\n", doubleSpace(b.Text))
	}
}

func doubleSpace(b []byte) []byte {
	return bytes.Replace(b, []byte("\n"), []byte("\n\n"), -1)
}
