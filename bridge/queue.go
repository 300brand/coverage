package bridge

import (
	"encoding/json"
	"git.300brand.com/coverage"
	"net/url"
	"time"
)

type Feed struct {
	QueueId uint64
	Id      uint64
	Url     string
}

type Report struct {
	QueueId    uint64
	Id         uint64
	DateBounds struct{ Start, End string }
	Phrases    struct{ Include, Exclude []string }
	Feeds      []string
	Matches    []string
	Pages      []string
	Summaries  []string
}

type Queue struct {
	LastId      uint64
	NewFeeds    []coverage.Feed
	RemoveFeeds []Feed
	Reports     []coverage.Report
}

type queueResponse struct {
	QueueId  uint64           `json:"id"`
	Class    string           `json:"class"`
	ObjectId uint64           `json:"object_id"`
	Data     *json.RawMessage `json:"data"`
}

const MySQLTime = "2006-01-02 15:04:05MST"

func GetQueue(LastId, Limit uint64) (q Queue, err error) {
	b := New()
	defer b.Close()

	q.LastId = LastId

	resp := make([]queueResponse, 0, Limit)
	err = b.Call("queue", []uint64{LastId, Limit}, &resp)
	if err != nil {
		return
	}

	for _, r := range resp {
		switch r.Class {
		case "CoverageReport":
			v := Report{QueueId: r.QueueId}
			if err = json.Unmarshal(*r.Data, &v); err != nil {
				return
			}
			q.Reports = append(q.Reports, convertReport(v))
		case "CoverageFeed":
			v := Feed{QueueId: r.QueueId, Id: r.ObjectId}
			if err = json.Unmarshal(*r.Data, &v.Url); err != nil {
				return
			}
			q.NewFeeds = append(q.NewFeeds, convertFeed(v))
		case "RemoveFeed":
			v := Feed{QueueId: r.QueueId, Id: r.ObjectId}
			if err = json.Unmarshal(*r.Data, &v.Url); err != nil {
				return
			}
			q.RemoveFeeds = append(q.RemoveFeeds, v)
		}
		q.LastId = r.QueueId
	}
	return
}

func convertFeed(in Feed) (out coverage.Feed) {
	out = *coverage.NewFeed()
	out.ObjectId = in.Id
	out.QueueId = in.QueueId
	out.URL, _ = url.Parse(in.Url)
	return

}

func convertReport(in Report) (out coverage.Report) {
	out = *coverage.NewReport()

	out.ObjectId = in.Id
	out.QueueId = in.QueueId

	tz := time.Now().Format("MST")
	out.DateBounds.Start, _ = time.Parse(MySQLTime, in.DateBounds.Start+tz)
	out.DateBounds.End, _ = time.Parse(MySQLTime, in.DateBounds.End+tz)

	out.Phrases = in.Phrases
	// Feed URLs
	out.Feeds = make([]*url.URL, len(in.Feeds))
	for i, u := range in.Feeds {
		out.Feeds[i], _ = url.Parse(u)
	}

	// Page/Previous Match URLs
	out.PreviousResults = make([]*url.URL, len(in.Pages))
	for i, u := range in.Pages {
		out.PreviousResults[i], _ = url.Parse(u)
	}
	return
}
