package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage/page/body"
	"io/ioutil"
	"os"
	"strings"
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
		fmt.Println(b.HTML)
	}
	if !noHTML && !noText {
		fmt.Println("\n========================================\n")
	}
	if !noText {
		fmt.Println(strings.Replace(b.Text, "\n", "\n\n", -1))
	}
}
