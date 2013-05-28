package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
)

func (s *StorageWriter) UpdateArticle(ri *skynet.RequestInfo, in *coverage.Article, out *coverage.Article) (err error) {
	return s.Mongo.UpdateArticle(in)
}

func (s *StorageWriter) UpdateFeed(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	return s.Mongo.UpdateFeed(in)
}

func (s *StorageWriter) UpdatePublication(ri *skynet.RequestInfo, in *coverage.Publication, out *coverage.Publication) (err error) {
	return s.Mongo.UpdatePublication(in)
}
