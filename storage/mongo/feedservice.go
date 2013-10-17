package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"github.com/jbaikge/logger"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type FeedService struct {
	m *Mongo
}

const (
	FeedCollection      = "Feeds"
	FeedQueueCollection = "FeedQ"
)

var _ service.FeedService = &FeedService{}

func NewFeedService(m *Mongo) *FeedService {
	return &FeedService{m: m}
}

func (s *FeedService) Update(f *coverage.Feed) error {
	f.Log.Service("mongo.FeedService")
	return s.m.UpdateFeed(f)
}

func (m *Mongo) FeedIds(urls []*url.URL) (ids []bson.ObjectId, err error) {
	logger.Trace.Printf("FeedIds: called")
	out := make([]struct {
		Id bson.ObjectId `bson:"_id"`
	}, 0, len(urls))
	err = m.C.Feeds.Find(bson.M{
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
	logger.Trace.Printf("GetFeed: called %+v", query)
	switch v := query.(type) {
	case bson.ObjectId:
		query = bson.M{"_id": v}
	}
	err = m.C.Feeds.Find(query).One(f)
	return
}

func (m *Mongo) GetOldestFeed(ignore []bson.ObjectId, f *coverage.Feed) (err error) {
	logger.Trace.Printf("GetOldestFeed: called")
	err = m.C.Feeds.Find(bson.M{
		"_id": bson.M{
			"$not": bson.M{
				"$in": ignore,
			},
		},
		"deleted": false,
	}).Sort("lastdownload").Limit(1).One(f)
	if err != nil {
		logger.Error.Printf("GetOldestFeed: %s", err)
	}
	return
}

func (m *Mongo) NextDownloadFeedId(thresh time.Time, id *bson.ObjectId) (err error) {
	n, err := m.C.FeedQ.Count()
	if err != nil {
		logger.Error.Printf("NextDownloadFeedId: %s", err)
		return
	}
	if n == 0 {
		logger.Trace.Print("NextDownloadFeedId: Queue too small, refilling")
		query := bson.M{
			"deleted": false,
		}
		sel := bson.M{"_id": 1}
		iter := m.C.Feeds.
			Find(query).
			Select(sel).
			Iter()

		var result struct {
			Id           bson.ObjectId `bson:"_id"`
			LastDownload time.Time
			Queue        int
		}

		imported := 0
		for iter.Next(&result) {
			if result.LastDownload.IsZero() {
				result.LastDownload = thresh
			}
			if err = m.C.FeedQ.Insert(result); err != nil {
				logger.Error.Printf("NextDownloadFeedId: Insertion error: %s", err)
				return
			}
			imported++
		}
		if err = iter.Close(); err != nil {
			logger.Error.Printf("NextDownloadFeedId: %s", err)
			return
		}
		logger.Trace.Printf("NextDownloadFeedId: Queue - inserted %d", imported)
	}

	query := bson.M{
		"queue":        0,
		"lastdownload": bson.M{"$lte": thresh},
	}
	change := mgo.Change{
		Remove:    true,
		ReturnNew: false,
		Update:    nil,
		Upsert:    false,
	}
	result := new(struct {
		Id bson.ObjectId `bson:"_id"`
	})
	info, err := m.C.FeedQ.Find(query).Limit(1).Apply(change, result)
	if err != nil {
		logger.Error.Printf("NextDownloadFeedId: %s", err)
		return
	}
	logger.Trace.Printf("NextDownloadFeedId: Updated %v", info.Updated)
	*id = result.Id
	return
}

func (m *Mongo) UpdateFeed(f *coverage.Feed) (err error) {
	logger.Trace.Printf("UpdateFeed: called")
	l := len(f.Articles)
	f.Articles = f.Articles[:0]
	defer func() { f.Articles = f.Articles[:l] }()

	f.Updated = time.Now()
	_, err = m.C.Feeds.UpsertId(f.ID, f)
	return
}
