package skytypes

import (
	"labix.org/v2/mgo/bson"
)

type ObjectIds struct {
	Ids []bson.ObjectId
}

type ObjectId struct {
	Id bson.ObjectId
}

type NullType struct{}

var Null = &NullType{}
