package main

import (
	"flag"
	"git.300brand.com/coverage/article/body"
	"git.300brand.com/coverage/article/lexer"
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/logger"
	"git.300brand.com/coverage/parser"
	"git.300brand.com/coverage/service"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

var config = struct {
	DBHost  string
	DBName  string
	ID      bson.ObjectId
	NewOnly bool
	Verbose bool
}{}

func init() {
	flag.BoolVar(&config.NewOnly, "new", false, "Only show new articles")
	flag.BoolVar(&config.Verbose, "v", false, "Verbose")
	flag.StringVar(&config.DBHost, "dbhost", "localhost", "Mongo server")
	flag.StringVar(&config.DBName, "dbname", "Coverage", "Mongo database name")
}

func main() {
	flag.Parse()

	if config.Verbose {
		logger.EnableLogging(os.Stdout)
	}

	if !bson.IsObjectIdHex(flag.Arg(0)) {
		log.Fatalf("Invalid bson.ObjectId: %s", flag.Arg(0))
	}
	config.ID = bson.ObjectIdHex(flag.Arg(0))

	m := mongo.New(config.DBHost, config.DBName)
	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	feed, err := m.GetFeed(config.ID)
	if err != nil {
		log.Fatalf("Could not find feed with ID: %s", config.ID.Hex())
	}

	log.Println("Downloading new feed")
	downloader.NewFeedService().Update(feed)

	log.Println("Saving data")
	mfs := mongo.NewFeedService(m)
	mfs.Update(feed)

	log.Println("Building article list")
	parser.NewFeedService().Update(feed)

	services := []service.ArticleService{
		downloader.NewArticleService(),
		body.NewArticleService(),
		lexer.NewArticleService(),
		mongo.NewArticleService(m),
	}
	for _, a := range feed.Articles {
		log.Printf("Processing %s", a.URL)
		for _, s := range services {
			if err := s.Update(a); err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Updating feed")
	mfs.Update(feed)
}
