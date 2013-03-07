package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/downloader"
	"io/ioutil"
	"net/http"
	"os"
)

var isFile bool

func init() {
	flag.BoolVar(&isFile, "f", false, "Argument is a file instead of URL")
}

func main() {
	flag.Parse()

	r, err := getResponse(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	a, err := coverage.NewArticle(&r)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	out, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Printf("%s\n", out)
}

func getResponse(path string) (r downloader.Response, err error) {
	if isFile {
		pwd, err := os.Getwd()
		if err != nil {
			return r, err
		}
		f, err := os.Open(path)
		if err != nil {
			return r, err
		}
		defer f.Close()
		if r.Body, err = ioutil.ReadAll(f); err != nil {
			return r, err
		}
		pathURL := fmt.Sprintf("file://%s/%s", pwd, path)
		r.Code = http.StatusOK
		r.OriginalURL = pathURL
		r.RealURL = pathURL
	} else {
		r, err = downloader.Fetch(path)
	}
	return
}
