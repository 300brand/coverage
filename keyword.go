package coverage

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Keyword struct {
	Hash      uint32 `bson:",minsize"`
	Keyword   string
	Date      time.Time
	ArticleId bson.ObjectId
	Published time.Time `bson:"-"`
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

func KeywordHash(s string) (hash uint32) {
	for i := 0; i < len(s); i++ {
		hash = hash*primeRK + uint32(s[i])
	}
	return
}
