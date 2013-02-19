package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"git.300brand.com/coverage/page/body"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body.CleanDOM(doc)

	if err = html.Render(os.Stderr, doc); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println()
}
