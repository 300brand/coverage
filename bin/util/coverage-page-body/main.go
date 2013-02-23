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

	var out []byte

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err = body.CleanHTML(in)
	if !cleanedOnly {
		out, err = body.GetBody(out)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("%s\n", out)
}
