package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
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
		mgo.Index{
			Key:        []string{"disabled"},
			Background: true,
			DropDups:   false,
			Sparse:     false,
			Unique:     false,
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

func (m *Mongo) GetFeed(query interface{}, f *coverage.Feed) (err error) {
	switch v := query.(type) {
	case bson.ObjectId:
		query = bson.M{"_id": v}
	}
	err = m.db.C(FeedCollection).Find(query).One(f)
	return
}

func (m *Mongo) GetOldestFeed(ignore []bson.ObjectId, f *coverage.Feed) (err error) {
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
	l := len(f.Articles)
	f.Articles = f.Articles[:0]
	defer func() { f.Articles = f.Articles[:l] }()

	f.Updated = time.Now()
	_, err = m.db.C(FeedCollection).UpsertId(f.ID, f)
	return
}
