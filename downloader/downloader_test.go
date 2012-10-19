package downloader

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "http://httpbin.org/response-headers?Content-Type=text/plain"
	expect := []byte(`{
  "Content-Length": "60",
  "Content-Type": "text/plain"
}`)
	r, err := Fetch(url)
	if err != nil {
		t.Error(err)
	}
	// Length check
	if len(r.Body) != len(expect) {
		t.Errorf("Expect: %s", expect)
		t.Errorf("Got:    %s", r.Body)
		return
	}
	for i, b := range expect {
		if r.Body[i] != b {
			t.Errorf("Invalid char `%s', expected `%s'", r.Body[i], b)
		}
	}
}

func TestRedirect(t *testing.T) {
	url := "http://httpbin.org/redirect/3"
	expect := "http://httpbin.org/get"
	r, err := Fetch(url)
	if err != nil {
		t.Error(err)
	}
	if r.RealURL != expect {
		t.Errorf("Expect: %s", expect)
		t.Errorf("Got:    %s", r.RealURL)
	}
}

func TestResponseCode(t *testing.T) {
	codes := []int{
		http.StatusOK,
		http.StatusNotFound,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
	}
	for _, code := range codes {
		r, err := Fetch(fmt.Sprintf("http://httpbin.org/status/%d", code))
		if err != nil {
			t.Error(err)
		}
		if r.Code != code {
			t.Errorf("Expect: %d", code)
			t.Errorf("Got:    %d", r.Code)
		}
	}
}
