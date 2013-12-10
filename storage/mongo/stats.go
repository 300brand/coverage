package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Stats struct {
	Publications collectionStat
	Feeds        collectionStat
	Articles     struct{ Tallied, Total int } // These should always be equal
	Queue        int
}

type collectionStat struct {
	Active, Inactive, Total int
}

func (s Stats) Hash() string {
	return fmt.Sprintf("%d.%d.%d.%d.%d.%d.%d.%d.%d",
		s.Publications.Active,
		s.Publications.Inactive,
		s.Publications.Total,
		s.Feeds.Active,
		s.Feeds.Inactive,
		s.Feeds.Total,
		s.Articles.Tallied,
		s.Articles.Total,
		s.Queue,
	)
}

// Based on the following:
//
// function report() {
//     o = {
//         date:     new Date,
//         pubs:     dbPublications.count(),
//         feeds:    dbFeeds.count(),
//         queue:    dbFeedQ.count(),
//         articles: dbArticles.count(),
//         tally:    dbPublications.find({numarticles: {$gt: 0}}, {numarticles:1, _id:0}).map(function (p) { return p.numarticles; }).reduce(function(prev, cur) { return prev + cur; })
//     };
//     o.delta = o.articles - o.tally;
//     print("As of " + o.date + " - P:" + o.pubs + " F:" + o.feeds + " Q:" + o.queue + " A/T:" + o.articles + "/" + o.tally + " (" + o.delta + ")");
// }
func (m *Mongo) GetStats(s *Stats) (err error) {
	if err = m.pubStat(s); err != nil {
		return
	}
	if err = m.feedStat(s); err != nil {
		return
	}
	if s.Articles.Total, err = m.C.Articles.Count(); err != nil {
		return
	}
	if s.Queue, err = m.C.FeedQ.Count(); err != nil {
		return
	}
	return
}

func (m *Mongo) feedStat(stats *Stats) (err error) {
	job := &mgo.MapReduce{
		Map: `
			function() {
				emit("Total", 1)
				if (this.deleted) {
					emit("Inactive", 1)
				} else {
					emit("Active", 1)
				}
			}`,
		Reduce: `
			function(key, values) {
				return Array.sum(values)
			}`,
	}
	var result []struct {
		Id    string `_id`
		Value int
	}
	if _, err = m.C.Feeds.Find(nil).MapReduce(job, &result); err != nil {
		return
	}
	for _, r := range result {
		switch r.Id {
		case "Total":
			stats.Feeds.Total = r.Value
		case "Inactive":
			stats.Feeds.Inactive = r.Value
		case "Active":
			stats.Feeds.Active = r.Value
		}
	}
	return
}

func (m *Mongo) pubStat(stats *Stats) (err error) {
	job := &mgo.MapReduce{
		Map: `
			function() {
				emit("Total", 1)
				emit("ArticleTally", this.numarticles)
				if (this.deleted) {
					emit("Inactive", 1)
				} else {
					emit("Active", 1)
				}
			}`,
		Reduce: `
			function(key, values) {
				return Array.sum(values)
			}`,
	}
	var result []struct {
		Id    string `_id`
		Value int
	}
	if _, err = m.C.Publications.Find(nil).MapReduce(job, &result); err != nil {
		return
	}
	for _, r := range result {
		switch r.Id {
		case "Total":
			stats.Publications.Total = r.Value
		case "Inactive":
			stats.Publications.Inactive = r.Value
		case "Active":
			stats.Publications.Active = r.Value
		case "ArticleTally":
			stats.Articles.Tallied = r.Value
		}
	}
	return
}
