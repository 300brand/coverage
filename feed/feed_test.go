package feed

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/feed/parser"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "git.300brand.com/coverage/feed/parser/atom"
	_ "git.300brand.com/coverage/feed/parser/rdf"
	_ "git.300brand.com/coverage/feed/parser/rss"
)

type test struct {
	Filename string
	Type     string
	Feed     *coverage.Feed
	Files    struct {
		Times string
		URLs  string
	}
}

var (
	feedTypes = []string{"atom", "rdf", "rss"}
	tests     = []test{}
)

func init() {
	for _, ext := range feedTypes {
		list, err := filepath.Glob("samples/*." + ext)
		if err != nil {
			log.Fatal(err)
		}
		for _, filename := range list {
			t := test{
				Filename: filename,
				Type:     ext,
			}
			t.Files.URLs = strings.Replace(filename, "."+ext, ".urls", 1)
			t.Files.Times = strings.Replace(filename, "."+ext, ".times", 1)

			t.Feed = coverage.NewFeed()
			if t.Feed.Content, err = ioutil.ReadFile(filename); err != nil {
				log.Fatalf("Parsing %s: %s", filename, err)
			}

			s := NewFeedService()
			if err = s.Update(t.Feed); err != nil {
				log.Fatalf("Parsing %s: %s", filename, err)
			}
			tests = append(tests, t)
		}
	}
}

func TestURLs(t *testing.T) {
	for _, test := range tests {
		if _, err := os.Stat(test.Files.URLs); err != nil {
			t.Logf("Skipping; no %s file found", test.Files.URLs)
			continue
		}

		b, err := ioutil.ReadFile(test.Files.URLs)
		if err != nil {
			t.Fatalf("Read error: %s", err)
		}

		urls := strings.Fields(string(b))
		if l := len(test.Feed.Articles); l != len(urls) {
			t.Fatalf("Invalid URL count\nExpect: %d\nGot: %d", len(urls), l)
		}

		for i, u := range urls {
			if aURL := test.Feed.Articles[i].URL; u != aURL.String() {
				t.Errorf("URL Mismatch\nExpect: %s\nGot: %s", u, aURL)
			}
		}
	}
}

func TestTimes(t *testing.T) {
	for _, test := range tests {
		if _, err := os.Stat(test.Files.Times); err != nil {
			t.Logf("Skipping; no %s file found", test.Files.Times)
			continue
		}

		b, err := ioutil.ReadFile(test.Files.Times)
		if err != nil {
			t.Fatalf("Read error: %s", err)
		}

		times := strings.FieldsFunc(string(b), func(r rune) bool { return r == '\n' })
		if l := len(test.Feed.Articles); l != len(times) {
			for _, a := range test.Feed.Articles {
				t.Logf("%s", a.Published)
			}
			t.Fatalf("Invalid time count\nExpect: %d\nGot: %d", len(times), l)
		}

		for i, exp := range times {
			if aTime := test.Feed.Articles[i].Published; exp != aTime.String() {
				t.Logf("Expected: %s; Got: %s", exp, aTime)
			}
		}
	}
}

func TestType(t *testing.T) {
	for _, test := range tests {
		typ, err := parser.Type(test.Feed.Content)
		if err != nil {
			t.Errorf("Decoder error [%s] test: %s", test.Filename, err)
		}
		if typ != test.Type {
			t.Errorf("Invalid type\nExpect: %s\nGot: %s", test.Type, typ)
		}
	}
}
