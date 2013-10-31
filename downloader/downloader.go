package downloader

import (
	"github.com/300brand/coverage/cleanurl"
	"io"
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

var MaxFileSize int64 = 2 * 1024 * 1024

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

	lr := io.LimitReader(resp.Body, MaxFileSize)
	r.Body, err = ioutil.ReadAll(lr)
	return
}
