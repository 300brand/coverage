package social

import (
	"encoding/json"
	"github.com/jbaikge/logger"
	"net/http"
	"net/url"
)

type Stats struct {
	Google   int `json:"GooglePlusOne"`
	Twitter  int
	LinkedIn int
	Facebook struct {
		Likes  int `json:"like_count"`
		Shares int `json:"share_count"`
	}
}

const apiURL = "http://api.sharedcount.com/"

func Fetch(u *url.URL, s *Stats) (err error) {
	logger.Trace.Printf("Fetch: called")
	return FetchString(u.String(), s)
}

func FetchString(u string, s *Stats) (err error) {
	logger.Trace.Printf("FetchString: called %s", u)
	api, err := url.Parse(apiURL)
	if err != nil {
		logger.Error.Printf("FetchString: %s", err)
		return
	}
	q := api.Query()
	q.Set("url", u)
	api.RawQuery = q.Encode()

	resp, err := http.Get(api.String())
	if err != nil {
		logger.Error.Printf("FetchString: %s", err)
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(s)
}
