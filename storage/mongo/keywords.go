package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"sync"
	"time"
)

const KeywordCollection = "Keywords3"

func (m *Mongo) ReduceKeywords(query interface{}) (info *mgo.MapReduceInfo, err error) {
	job := &mgo.MapReduce{
		// Map: `function() {
		// 	for (var i in this.text.words.keywords) {
		// 		var key = this.text.words.keywords[i]
		// 		var value = {
		// 			count:    1,
		// 			articles: [
		// 				{
		// 					id:        this._id,
		// 					published: this.published
		// 				}
		// 			]
		// 		}
		// 		emit(key, value)
		// 	}
		// }`,
		// Reduce: `function(key, values) {
		// 	var list = { count: 0, articles: [] }
		// 	for (var i in values) {
		// 		list.articles = values[i].articles.concat(list.articles)
		// 	}
		// 	list.count = list.articles.length
		// 	return list
		// }`,
		Map: `function() {
			var d
			if (this.published > new Date(0)) {
				d = this.published
			} else {
				d = this.added
			}
			var key = {
				keyword: "",
				date: new Date(d.getFullYear(), d.getMonth(), d.getDate())
			}
			var value = {
				count: 1,
				articles: [this._id]
			}
			for (var i in this.text.words.keywords) {
				key.keyword = this.text.words.keywords[i]
				emit(key, value)
			}
		}`,
		Reduce: `function(key, values) {
			var list = { count: 0, articles: [] }
			for (var i in values) {
				list.articles = values[i].articles.concat(list.articles)
			}
			list.count = list.articles.length
			return list
		}`,
		// Map: `function() {
		// 	var d
		// 	if (this.published > new Date(0)) {
		// 		d = this.published
		// 	} else {
		// 		d = this.added
		// 	}
		// 	var dStr = d.getFullYear()*10000 + d.getMonth()*100 + d.getDate()
		// 	var value = {}
		// 	value[dStr] = [this._id]
		// 	var keywords = this.text.words.keywords
		// 	for (var i in keywords) {
		// 		emit(keywords[i], value)
		// 	}
		// }`,
		// Reduce: `function(key, values) {
		// 	var value = values[0]
		// 	for (var i = 1; i < values.length; i++) {
		// 		for (var d in values[i]) {
		// 			value[d]
		// 		}
		// 		list.articles = values[i].articles.concat(list.articles)
		// 	}
		// 	list.count = list.articles.length
		// 	return list
		// }`,
		Out:     bson.M{"reduce": KeywordCollection},
		Verbose: true,
	}
	return m.db.C(ArticleCollection).Find(query).MapReduce(job, nil)
}

func (m *Mongo) KeywordSearch(keywords []string, from, to time.Time, kwChan chan coverage.Keyword) (err error) {
	hashes := make([]uint32, len(keywords))
	for i, w := range keywords {
		hashes[i] = coverage.KeywordHash(w)
	}

	query := bson.M{}
	if from.After(time.Time{}) || to.After(time.Time{}) {
		query["date"] = bson.M{
			"$gte": from.Truncate(24 * time.Hour),
			"$lte": to.Truncate(24 * time.Hour),
		}
	}

	var wg sync.WaitGroup
	for i, word := range keywords {
		log.Printf("mongo.KeywordSearch: Starting go func(%s)", word)
		wg.Add(1)
		go func(i int) {
			q := bson.M{
				"hash": hashes[i],
			}
			log.Printf("mongo.KeywordSearch: Query %v", q)
			iter := m.kdb.C(KeywordCollection).Find(q).Iter()
			kw := &coverage.Keyword{}
			for iter.Next(kw) {
				kwChan <- *kw
			}
			if err = iter.Close(); err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(i)
	}
	log.Printf("Waiting for sends to finish")
	wg.Wait()
	close(kwChan)
	/*
			find := bson.M{
				"_id:": bson.M{
					"$in": keywords,
				},
			}
			result := &struct {
				Word string `bson:"_id"`
				Ids  []struct {
					Id        bson.ObjectId
					Published time.Time
				}
			}{}
			ids = make([]bson.ObjectId, 0, 1024)

			iter := m.db.C(KeywordCollection).Find(find).Iter()
			for iter.Next(result) {
				log.Printf("%+v", result)
				for _, id := range result.Ids {
					if id.Published.After(from) && id.Published.Before(to) {
						ids = append(ids, id.Id)
					}
				}
			}
			err = iter.Close()
			return

		aggregate := []bson.M{
			// Pull all article lists with matching keywords
			bson.M{
				"$match": bson.M{
					"_id": bson.M{
						"$in": keywords,
					},
				},
			},
			// Expand the articles element and build a new document for each one
			bson.M{"$unwind": "$value.articles"},
			// Filter each by date bounds
			bson.M{
				"$match": bson.M{
					"value.articles.published": bson.M{
						"$gte": from,
						"$lte": to,
					},
				},
			},
			// Reassemble documents by ID and count the word matches
			bson.M{
				"$group": bson.M{
					"_id": "$value.articles.id",
					"words": bson.M{
						"$push": "$_id",
					},
					"count": bson.M{
						"$sum": 1,
					},
				},
			},
			// Match documents that matched all len(keywords)
			bson.M{
				"$match": bson.M{
					"count": len(keywords),
				},
			},
		}

		type row struct {
			Id bson.ObjectId `bson:"_id"`
		}
		result := &struct {
			Ids []row
		}{}
		err = m.db.C(KeywordCollection).Pipe(aggregate).One(&result)
		ids = result.Ids
	*/
	return
}
