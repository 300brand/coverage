package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const KeywordCollection = "Keywords"

func (m *Mongo) ReduceKeywords(query interface{}) (info *mgo.MapReduceInfo, err error) {
	job := &mgo.MapReduce{
		Map: `function() {
			for (var i in this.text.words.keywords) {
				var key = this.text.words.keywords[i]
				var value = {
					count:    1,
					articles: [
						{
							id:        this._id,
							published: this.published
						}
					]
				}
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
		Out:     bson.M{"reduce": KeywordCollection},
		Verbose: true,
	}
	return m.db.C(ArticleCollection).Find(query).MapReduce(job, nil)
}

func (m *Mongo) KeywordArticleIds(keywords []string, from, to time.Time) (ids []bson.ObjectId, err error) {
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
		// Reassemble the unique documents
		bson.M{
			"$group": bson.M{
				// _id MUST be specified, using zero to group all the articles together
				"_id": "0",
				"articles": bson.M{
					// $addToSet uniques the article ids
					"$addToSet": "$value.articles.id",
				},
			},
		},
	}

	result := &struct {
		Ids []bson.ObjectId `bson:"articles"`
	}{}
	err = m.db.C(KeywordCollection).Pipe(aggregate).One(&result)
	ids = result.Ids
	return
}
