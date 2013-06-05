package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
)

func (s *StorageWriter) SaveArticle(ri *skynet.RequestInfo, in *coverage.Article, out *coverage.Article) (err error) {
	if err = s.Mongo.UpdateArticle(in); err != nil {
		return
	}
	*out = *in
	return
}

func (s *StorageWriter) SaveFeed(ri *skynet.RequestInfo, in *coverage.Feed, out *coverage.Feed) (err error) {
	if err = s.Mongo.UpdateFeed(in); err != nil {
		return
	}
	*out = *in
	return
}

func (s *StorageWriter) SavePublication(ri *skynet.RequestInfo, in *coverage.Publication, out *coverage.Publication) (err error) {
	if err = s.Mongo.UpdatePublication(in); err != nil {
		return
	}
	*out = *in
	return
}
