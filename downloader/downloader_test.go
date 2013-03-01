package downloader

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	Success = "success"
)

type status int

var codes = []int{
	http.StatusOK,
	http.StatusNotFound,
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
}

func (s status) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(int(s))
}

func server() *httptest.Server {
	mux := http.NewServeMux()
	// Default response
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, Success)
	})
	// Redirect
	mux.HandleFunc("/redirect", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	})
	// Codes
	for _, code := range codes {
		mux.Handle(fmt.Sprintf("/%d", code), status(code))
	}
	return httptest.NewServer(mux)
}

func TestDownload(t *testing.T) {
	s := server()
	defer s.Close()

	r, err := Fetch(s.URL)
	if err != nil {
		t.Error(err)
	}
	if string(r.Body) != Success {
		t.Errorf("Expect: %s", Success)
		t.Errorf("Got:    %s", r.Body)
	}
}

func TestRealURL(t *testing.T) {
	s := server()
	defer s.Close()

	r, err := Fetch(s.URL)
	if err != nil {
		t.Error(err)
	}
	if r.RealURL == "" {
		t.Error("RealURL not set")
	}
}

func TestRedirect(t *testing.T) {
	s := server()
	defer s.Close()

	url := s.URL + "/redirect"
	expect := s.URL + "/"
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
	s := server()
	defer s.Close()

	for _, code := range codes {
		r, err := Fetch(fmt.Sprintf("%s/%d", s.URL, code))
		if err != nil {
			t.Error(err)
		}
		if r.Code != code {
			t.Errorf("Expect: %d", code)
			t.Errorf("Got:    %d", r.Code)
		}
	}
}
