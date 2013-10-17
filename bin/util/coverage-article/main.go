package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/article/body"
	"github.com/300brand/coverage/article/lexer"
	"github.com/300brand/coverage/downloader"
	"github.com/300brand/coverage/service"
	"github.com/300brand/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"net/url"
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
	a.URL = fixURL(flag.Arg(0))

	services := []service.ArticleService{
		downloader.NewArticleService(),
		body.NewArticleService(),
		lexer.NewArticleService(),
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

func fixURL(s string) *url.URL {
	if isFile {
		if s[0] != '/' {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			s = fmt.Sprintf("%s%s%s", wd, string(os.PathSeparator), s)
		}
		s = "file://" + s
	}
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return u
}
