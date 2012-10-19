package downloader

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Body        []byte
	Code        int
	OriginalURL string
	RealURL     string
}

func Fetch(url string) (r Response, err error) {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) (err error) {
			// Add checks to remove ?rss and other params
			r.RealURL = req.URL.String()
			return
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	r.OriginalURL = url
	r.Code = resp.StatusCode
	r.Body, err = ioutil.ReadAll(resp.Body)
	return
}
