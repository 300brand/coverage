package bridge

import (
	"encoding/json"
	"git.300brand.com/coverage"
	"net/url"
	"time"
)

type Feed struct {
	Id  uint64
	Url string
}

type Report struct {
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
	NewFeeds    []Feed
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

func GetQueue(LastId, Limit int) (q Queue, err error) {
	b := New()
	defer b.Close()

	resp := make([]queueResponse, 0, Limit)
	err = b.Call("queue", []int{LastId, Limit}, &resp)
	if err != nil {
		return
	}

	for _, r := range resp {
		switch r.Class {
		case "CoverageReport":
			v := Report{}
			if err = json.Unmarshal(*r.Data, &v); err != nil {
				return
			}
			q.Reports = append(q.Reports, convertReport(v))
		case "CoverageFeed":
			v := Feed{Id: r.ObjectId}
			if err = json.Unmarshal(*r.Data, &v.Url); err != nil {
				return
			}
			q.NewFeeds = append(q.NewFeeds, v)
		case "RemoveFeed":
			v := Feed{Id: r.ObjectId}
			if err = json.Unmarshal(*r.Data, &v.Url); err != nil {
				return
			}
			q.RemoveFeeds = append(q.RemoveFeeds, v)
		}
		q.LastId = r.QueueId
	}
	return
}

func convertReport(in Report) (out coverage.Report) {
	out.ObjectId = in.Id
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
