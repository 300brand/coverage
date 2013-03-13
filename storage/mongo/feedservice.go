package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type FeedService struct {
	m *Mongo
}

const FeedCollection = "Feeds"

var _ service.FeedService = &FeedService{}

func init() {
	indexes[FeedCollection] = []mgo.Index{
		mgo.Index{
			Key:        []string{"url"},
			Background: true,
			DropDups:   true,
			Sparse:     false,
			Unique:     true,
		},
	}
}

func NewFeedService(m *Mongo) *FeedService {
	return &FeedService{m: m}
}

func (s *FeedService) Update(f *coverage.Feed) error {
	f.Log.Service("mongo.FeedService")
	return s.m.UpdateFeed(f)
}

func (m *Mongo) GetFeed(query interface{}) (f *coverage.Feed, err error) {
	switch v := query.(type) {
	case bson.ObjectId:
		query = bson.M{"_id": v}
	}
	f = &coverage.Feed{}
	err = m.db.C(FeedCollection).Find(query).One(f)
	return
}

func (m *Mongo) UpdateFeed(f *coverage.Feed) (err error) {
	f.Log.Debug("mongo.UpdateFeed")
	_, err = m.db.C(FeedCollection).UpsertId(f.ID, f)
	if err != nil {
		return
	}
	for _, file := range f.Files() {
		if err = m.storeFile(FeedCollection, &file); err != nil {
			return
		}
	}
	return
}
