package skytypes

import (
	"labix.org/v2/mgo/bson"
)

type NullType struct{}

type ObjectIds struct {
	Ids []bson.ObjectId
}

type ObjectId struct {
	Id bson.ObjectId
}

var Null = &NullType{}
