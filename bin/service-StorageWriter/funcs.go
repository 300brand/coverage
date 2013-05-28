package main

import (
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
)

func (s *StorageWriter) UpdateArticle(ri *skynet.RequestInfo, a *coverage.Article, out *coverage.Article) (err error) {
	return s.Mongo.UpdateArticle(a)
}

func (s *StorageWriter) UpdateFeed(ri *skynet.RequestInfo, f *coverage.Feed, out *coverage.Feed) (err error) {
	return s.Mongo.UpdateFeed(f)
}

func (s *StorageWriter) UpdatePublication(ri *skynet.RequestInfo, p *coverage.Publication, out *coverage.Publication) (err error) {
	return s.Mongo.UpdatePublication(p)
}
