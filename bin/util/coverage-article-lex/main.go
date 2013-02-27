package main

import (
	"fmt"
	"git.300brand.com/coverage/article/lexer"
	"io/ioutil"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	
}
