package main

import (
	"flag"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

var dbHost, dbName string

func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "MongoDB host")
	flag.StringVar(&dbName, "dbName", "Coverage", "MongoDB database")
}

func main() {
	flag.Parse()

	db, err := mgo.Dial(dbHost)
	if err != nil {
		log.Fatal(err)
	}

	c := db.DB(dbName).C(mongo.ArticleCollection)

	job := &mgo.MapReduce{
		Map: `function() {
			var r = db.Feed.find({_id: this.feedid}).
			emit(this.feedid, 1)
		}`,
		Reduce: `function(key, values) {
			return Array.sum(values)
		}`,
	}

	var result []struct {
		Id    bson.ObjectId `bson:"_id"`
		Value int
	}

	if _, err := c.Find(nil).MapReduce(job, &result); err != nil {
		log.Fatal(err)
	}

	for i, r := range result {
		log.Printf("[%d] %s %d", i, r.Id.Hex(), r.Value)
	}
}
