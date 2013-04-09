package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/url"
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

func (m *Mongo) FeedIds(urls []*url.URL) (ids []bson.ObjectId, err error) {
	out := make([]struct {
		Id bson.ObjectId `bson:"_id"`
	}, 0, len(urls))
	err = m.db.C(FeedCollection).Find(bson.M{
		"url": bson.M{
			"$in": urls,
		},
	}).Select(bson.M{"_id": 1}).All(&out)
	if err != nil {
		return
	}
	ids = make([]bson.ObjectId, len(out))
	for i := range out {
		ids[i] = out[i].Id
	}
	return
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

func (m *Mongo) GetOldestFeed(ignore []bson.ObjectId) (f *coverage.Feed, err error) {
	f = &coverage.Feed{}
	err = m.db.C(FeedCollection).Find(bson.M{
		"_id": bson.M{
			"$not": bson.M{
				"$in": ignore,
			},
		},
	}).Sort("lastdownload").Limit(1).One(f)
	return
}

func (m *Mongo) UpdateFeed(f *coverage.Feed) (err error) {
	if err = m.EnsureIndexSet(FeedCollection, indexes[FeedCollection]); err != nil {
		return
	}

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
