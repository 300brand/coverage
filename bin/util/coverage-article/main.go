package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/article/body"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/service"
	"os"
)

var isFile bool

func init() {
	flag.BoolVar(&isFile, "f", false, "Argument is a file instead of URL")
}

func main() {
	flag.Parse()

	a := coverage.NewArticle()
	url := fixURL(flag.Arg(0))
	services := []service.Service{
		downloader.NewService(url),
		body.NewService(),
	}

	for i, s := range services {
		if err := s.Update(a); err != nil {
			fmt.Printf("[%d] service error: %s\n", i, err)
			os.Exit(2)
		}
	}

	out, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Printf("%s\n", out)
}

func fixURL(url string) string {
	if isFile {
		if url[0] != '/' {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			url = fmt.Sprintf("%s%s%s", wd, string(os.PathSeparator), url)
		}
		url = "file://" + url
	}
	return url
}
