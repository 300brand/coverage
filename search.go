package coverage

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Search struct {
	Id       bson.ObjectId `bson:"_id"`
	Notify   string
	Q        string
	DaysLeft int
	Results  int
	Start    time.Time
	Complete *time.Time
	Dates    struct {
		Start, End time.Time
	}
	Articles []bson.ObjectId
}

type SearchResult struct {
	SearchId  bson.ObjectId
	ArticleId bson.ObjectId
}
