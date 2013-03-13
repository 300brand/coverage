package downloader

import (
	"git.300brand.com/coverage/cleanurl"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Body        []byte
	Code        int
	OriginalURL string
	RealURL     string
}

func Fetch(URL string) (r Response, err error) {
	r.OriginalURL = URL

	// Add support for the file protocol
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			r.RealURL = cleanurl.Clean(req.URL).String()
			return nil, nil
		},
	}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	r.Code = resp.StatusCode
	r.Body, err = ioutil.ReadAll(resp.Body)
	return
}
