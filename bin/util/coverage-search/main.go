package main

import (
	"flag"
	"git.300brand.com/coverage/storage/mongo"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

var dbHost, dbName string

func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "MongoDB host")
	flag.StringVar(&dbName, "dbName", "Coverage", "MongoDB database")
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	flag.Parse()

	db, err := mgo.Dial(dbHost)
	if err != nil {
		log.Fatal(err)
	}

	c := db.DB(dbName).C(mongo.ArticleCollection)

	//var pubs []struct {
	//	Id    bson.ObjectId `bson:"_id"`
	//	Title string
	//	Url   string
	//	Feeds []bson.ObjectId
	//}
	//db.DB(dbName).C(mongo.PublicationCollection).Find(nil).Select(bson.M{
	//	"title": 1,
	//	"url":   1,
	//	"feeds": 1,
	//}).All(&pubs)

	job := &mgo.MapReduce{
		Map: `function() {
			var value = {
				article: {
					feedid: this.feedid,
					title:  this.title,
					url:    this.url
				},
				sensitive: false,
				index: 0
			}
			var _id = this._id
			var qLen = q.length
			var words = this.words
			var wordLen = words.length - qLen
			// Don't bother if we don't have enough words
			if (wordLen < 0) {
				return
			}

			var indexes = function(a, w) {
				var list = []
				for (var i = -1; (i = a.indexOf(w, i)) > -1; i++) {
					list.push(i)
				}
				emit(_id, {len: a.length, list: list})
				return list
			}

			indexes(words, q[0]).forEach(function(i) {
				for (var j = 1; j < qLen; j++) {
					if (words[i+j].word != q[j]) {
						return
					}
				}
				value.sensitive = true
				value.index = i
				emit(_id, value)
			})

			//var iwords = []
			//var iq = []
			//words.forEach(function(w) {
			//	iwords.push(w.word.toLowerCase())
			//})
			//q.forEach(function(w) {
			//	iq.push(w.toLowerCase())
			//})
			//var idx = 0
			//while (Array.indexOf(words, q[0], idx))
//
//
			//var cs, ci
			//for (var i = 0; i < wordLen; i++) {
			//	cs = 0; ci = 0
			//	for (var j = 0; j < qLen; j++) {
			//		if (words[i+j].word == q[j]) {
			//			cs++
			//		}
			//	}
				//for (var j = 0; j < qLen; j++) {
				//	if (iwords[i+j] == iq[j]) {
				//		ci++
				//	}
				//}
			//	if (cs == qLen) {
			//		value.sensitive = [i]
			//		emit(this._id, value)
			//	}
				//if (ci == qLen) {
				//	value.insensitive = [i]
				//	//emit(this._id, value)
				//}
			//}
		}`,
		Reduce: `function(key, values) {
			//var out = values[0]
			//for (var i = 1; i < values.length; i++) {
			//	out.score += values[i].score
			//}
			//if (value.caseSensitive > 0 || value.caseInsensitive > 0) {
			//	value.score = ((value.caseSensitive * 2 + value.caseInsensitive) / value.positions[0]) * value.positions.length
			//}
			return values[0]
		}`,
		Scope: struct {
			Q []string
		}{
			Q: flag.Args(),
		},
		Out:     bson.M{"replace": "SearchResults"},
		Verbose: true,
	}

	var result []struct {
		Id    bson.ObjectId `bson:"_id"`
		Value struct {
			Counter int
		}
	}
	log.Println("Querying...")
	info, err := c.Find(bson.M{
		"words.word": bson.M{
			"$in": flag.Args(),
		},
	}).MapReduce(job, nil)
	if err != nil {
		log.Fatalf("MapReduce: %s", err)
	}
	log.Printf("Number of documents mapped %d", info.InputCount)
	log.Printf("Number of times reduce called emit %d", info.EmitCount)
	log.Printf("Number of documents in resulting collection %d", info.OutputCount)
	log.Printf("Output database: %s", info.Database)
	log.Printf("Output collection: %s", info.Collection)
	log.Printf("Time to run the job: %s", time.Duration(info.Time))

	for i, r := range result {
		log.Printf("[%d] %s %v", i, r.Id.Hex(), r.Value)
	}
}
