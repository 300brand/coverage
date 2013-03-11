package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/article/body"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/service"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"os"
)

var (
	isFile bool
	feedId string
)

func init() {
	flag.BoolVar(&isFile, "f", false, "Argument is a file instead of URL")
	flag.StringVar(&feedId, "feedId", "", "If supplied, saves article to the database under the supplied Feed ID")
}

func main() {
	flag.Parse()

	a := coverage.NewArticle()
	url := fixURL(flag.Arg(0))
	services := []service.ArticleService{
		downloader.NewArticleService(url),
		body.NewArticleService(),
	}

	if feedId != "" {
		m := mongo.New("localhost", "Coverage")
		m.Connect()
		defer m.Close()
		services = append(services, mongo.NewArticleService(m))
		a.FeedId = bson.ObjectIdHex(feedId)
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
