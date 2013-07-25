package coverage

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Keyword struct {
	Id       KeywordId `bson:"_id"`
	Articles []bson.ObjectId
}

type KeywordId struct {
	Keyword string
	Date    time.Time
}
