package downloader

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response")
	}))
	defer server.Close()
	expect := "Response"
	r, err := Fetch(server.URL)
	if err != nil {
		t.Error(err)
	}
	if string(r.Body) != expect {
		t.Errorf("Expect: %s", expect)
		t.Errorf("Got:    %s", r.Body)
	}
}

func TestRealURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response")
	}))
	defer server.Close()
	r, err := Fetch(server.URL)
	if err != nil {
		t.Error(err)
	}
	if r.RealURL == "" {
		t.Error("RealURL not set")
	}
}

func TestRedirect(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response")
	}))
	defer server.Close()

	url := "http://httpbin.org/redirect/3"
	expect := "http://httpbin.org/get"
	r, err := Fetch(server.URL)
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
