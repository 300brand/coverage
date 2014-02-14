package coverage

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type GroupSearch struct {
	Id     bson.ObjectId `bson:"_id"`
	Type   string        // Reliable way to determine the search type when sent over the line
	Notify struct {
		Done   string
		Social string
	}
	Start     time.Time
	Complete  *time.Time
	SearchIds []bson.ObjectId
	Searches  []Search `bson:-`
}

type Search struct {
	Id     bson.ObjectId `bson:"_id"`
	Type   string        // Reliable way to determine the search type when sent over the line
	Notify struct {
		Done   string
		Social string
	}
	Q        string
	Label    string
	DaysLeft int
	Results  int
	Start    time.Time
	Complete *time.Time
	Dates    struct {
		Start, End time.Time
	}
	Articles       []bson.ObjectId
	PublicationIds []bson.ObjectId
}

type SearchResult struct {
	SearchId  bson.ObjectId
	ArticleId bson.ObjectId
}

func NewSearch() *Search {
	return &Search{
		Id:   bson.NewObjectId(),
		Type: "Search",
	}
}

func NewGroupSearch() *GroupSearch {
	return &GroupSearch{
		Id:   bson.NewObjectId(),
		Type: "GroupSearch",
	}
}
