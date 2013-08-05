package skytypes

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
	"time"
)

type ClockCommand struct {
	Command string
	Tick    time.Duration
}

type ClockResult struct {
	Message string
}

type DateSearch struct {
	Id    bson.ObjectId
	Date  time.Time
	Query string
}

type NullType struct{}

type ObjectIds struct {
	Ids []bson.ObjectId
}

type ObjectId struct {
	Id bson.ObjectId
}

type SearchQuery struct {
	Q      string
	Notify string
	Dates  struct {
		Start, End time.Time
	}
}

type SearchQueryResponse struct {
	Id    bson.ObjectId
	Start time.Time
}

type SearchStatus struct {
	Id bson.ObjectId
}

type SearchResults struct {
	Id        bson.ObjectId
	Ready     bool
	Completed time.Time
	Articles  []coverage.Article
}

var Null = &NullType{}
