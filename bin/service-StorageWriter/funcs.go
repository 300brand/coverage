package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/mgo/bson"
	"github.com/skynetservices/skynet"
)

func (s *StorageWriter) EnqueueNextFeed(ri *skynet.RequestInfo, in interface{}, out interface{}) (err error) {
	r := c.GetService("StorageReader", "", "", "")
	ids := make([]bson.ObjectId, 0, mongo.DefaultCapacity)
	if err = r.Send(ri, "QueueFeedIds", nil, ids); err != nil {
		return
	}
	feed := &coverage.Feed{}
	if err = r.Send(ri, "OldestFeed", ids, feed); err != nil {
		return
	}
	return
}

func (s *StorageWriter) SaveArticle(ri *skynet.RequestInfo, in *coverage.Article, out *coverage.Article) (err error) {
	return s.Mongo.UpdateArticle(in)
}

func (s *StorageWriter) SaveFeed(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	*out = *in
	return s.Mongo.UpdateFeed(in)
}

func (s *StorageWriter) SavePublication(ri *skynet.RequestInfo, in *coverage.Publication, out *coverage.Publication) (err error) {
	*out = *in
	return s.Mongo.UpdatePublication(in)
}
