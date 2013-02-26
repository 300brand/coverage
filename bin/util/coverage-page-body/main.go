package main

import (
	"flag"
	"fmt"
	"git.300brand.com/coverage/page/body"
	"io/ioutil"
	"os"
)

var (
	cleanedOnly bool
)

func init() {
	flag.BoolVar(&cleanedOnly, "clean", false, "Only show the cleaned HTML used before extracing body content")
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
	fmt.Printf("Body: %v\n", b)
}
