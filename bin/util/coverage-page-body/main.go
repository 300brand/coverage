package main

import (
	//"code.google.com/p/go.net/html"
	"fmt"
	"git.300brand.com/coverage/page/body"
	"io/ioutil"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := body.GetBody(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("Body:")
	fmt.Println(b)
}
