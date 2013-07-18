package coverage

import (
	"github.com/jbaikge/rabinkarp"
	"labix.org/v2/mgo/bson"
	"sync"
	"time"
)

type Keyword struct {
	Id        KeywordId     `bson:"_id"`
	ArticleId bson.ObjectId `bson:"-"`
	Articles  []bson.ObjectId
}

type KeywordId struct {
	Keyword string
	Date    time.Time
	//Hash uint32 `bson:",minsize"`
}

var (
	keywordHasher = rabinkarp.New32()
	kHMutex       sync.Mutex
)

func KeywordHash(s string) (hash uint32) {
	kHMutex.Lock()
	defer kHMutex.Unlock()
	keywordHasher.Write([]byte(s))
	hash = keywordHasher.Sum32()
	keywordHasher.Reset()
	return
}
