package main

import (
	"database/sql"
	"flag"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/article/body"
	"github.com/300brand/coverage/article/lexer"
	"github.com/300brand/coverage/service"
	"github.com/300brand/coverage/storage/mongo"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Article struct {
	Id        uint64
	FeedId    uint64
	Title     string
	Url       string
	Published string
	Added     string
}

var conn struct {
	MySQL       *sql.DB
	Mongo       *mongo.Mongo
	ArticleStmt *sql.Stmt
}

const timeLayout = "2006-01-02 15:04:05"

var (
	batchSize uint64
	once      bool
	start     uint64
	loc       *time.Location
	path      string
	services  []service.ArticleService
)

func init() {
	var err error

	flag.BoolVar(&once, "once", false, "Only run one batch")
	flag.StringVar(&path, "path", "/media/bodies", "Path to article bodies")
	flag.Uint64Var(&batchSize, "batch", 100, "Batch size")
	flag.Uint64Var(&start, "start", 0, "Page ID to start with")

	log.SetFlags(log.Lmicroseconds)

	if loc, err = time.LoadLocation("America/New_York"); err != nil {
		log.Fatal(err)
	}

	conn.Mongo = mongo.New("localhost", "Coverage")
	if err = conn.Mongo.Connect(); err != nil {
		log.Fatal(err)
	}

	if err = conn.Mongo.EnsureIndexes(); err != nil {
		log.Fatal(err)
	}

	services = []service.ArticleService{
		body.NewArticleService(),
		lexer.NewArticleService(),
		mongo.NewArticleService(conn.Mongo),
	}

	fServices = []service.FeedService{
		mongo.NewFeedService(conn.Mongo),
	}

	if conn.MySQL, err = sql.Open("mysql", "root:@/haha?charset=utf8"); err != nil {
		log.Fatal(err)
	}

	conn.ArticleStmt, err = conn.MySQL.Prepare(`
		SELECT id, parent_id, IFNULL(title, ''), url, IFNULL(published, ''), _added
		FROM Page
		WHERE id > ?
			AND parent_id IS NOT NULL
		ORDER BY id
		LIMIT ?
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	processed := make(chan interface{}, batchSize)
	batchAdvance := make(chan bool)

	defer conn.MySQL.Close()
	defer conn.Mongo.Close()
	defer close(processed)

	if err := ImportFeeds(); err != nil {
		log.Println("Already imported")
	}

	go func(ch chan interface{}) {
		var i uint64 = 0
		for c := range ch {
			switch t := c.(type) {
			case error:
				log.Printf("[%06d] Error: %s", i, t)
			case *coverage.Article:
				log.Printf("[%06d] %s %s %s", i, t.ID.Hex(), t.FeedId.Hex(), t.Added)
			default:
				log.Printf("Unknown? %+v", t)
			}
			i++
			if i%batchSize == 0 {
				batchAdvance <- true
			}
		}
	}(processed)

	batch := make([]Article, batchSize)
	for {
		log.Printf("Processing batch of %d starting at %d", batchSize, start)
		n, err := GetBatch(start, batch)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Batch size: %d", n)
		if n == 0 {
			break
		}
		start = ProcessBatch(batch[:n], processed)
		if once {
			break
		}
		<-batchAdvance
	}

}
