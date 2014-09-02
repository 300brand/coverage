package downloader

import (
	"errors"
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

	redirectPolicyFunc := func(req *http.Request, via []*http.Request) error {
		if len(via) > 13 {
			return errors.New("Fetch URL Error: Exceeded limit of 14 redirects")
		}
		return nil
	}

	client := &http.Client{
		Transport:     transport,
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return
	}
	// Fix for www.healthmgttech.com:
	req.Header.Add("Accept-Encoding", "identity")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	r.Code = resp.StatusCode
	// Pull the remote URL out of the subscription service (if in use)
	if xRemoteUrl := resp.Header.Get("X-Remote-Url"); xRemoteUrl != "" {
		r.RealURL = xRemoteUrl
	}

	lr := io.LimitReader(resp.Body, MaxFileSize)
	r.Body, err = ioutil.ReadAll(lr)
	return
}
