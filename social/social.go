package social

import (
	"encoding/json"
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

func Fetch(u *url.URL) (s Stats, err error) {
	return FetchString(u.String())
}

func FetchString(u string) (s Stats, err error) {
	api, err := url.Parse(apiURL)
	if err != nil {
		return
	}
	q := api.Query()
	q.Set("url", u)
	api.RawQuery = q.Encode()

	resp, err := http.Get(api.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&s)
	return
}
