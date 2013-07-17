package coverage

import (
	"hash/fnv"
	"labix.org/v2/mgo/bson"
	"sync"
	"time"
)

type Keyword struct {
	Hash      uint32 `bson:",minsize"`
	Date      time.Time
	Keyword   string
	ArticleId bson.ObjectId `bson:"-"`
	Articles  []bson.ObjectId
}

var (
	keywordHasher = fnv.New32a()
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
