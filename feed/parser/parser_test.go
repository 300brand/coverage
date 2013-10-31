package parser

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "github.com/300brand/coverage/feed/parser/atom"
	_ "github.com/300brand/coverage/feed/parser/rdf"
	_ "github.com/300brand/coverage/feed/parser/rss"
)

type feedTest struct {
	Data  []byte
	Name  string
	Path  string
	Times []time.Time
	Type  string
	URLs  []*url.URL
}

func TestParse(t *testing.T) {
	for name, test := range getFiles(t) {
		_, err := Parse(test.Data)
		if err != nil {
			t.Errorf("%s: %s", name, err)
		}
	}
}

func TestParseType(t *testing.T) {
	for name, test := range getFiles(t) {
		_, err := ParseType(test.Data, test.Type)
		if err != nil {
			t.Errorf("%s: %s", name, err)
		}
	}
}

func TestType(t *testing.T) {
	for name, test := range getFiles(t) {
		if got, err := Type(test.Data); err != nil {
			t.Errorf("%s: Got error: %s", name, err)
		} else if got != test.Type {
			t.Errorf("%s: Got: %s Exp: %s", name, got, test.Type)
		}
	}
}

func getFiles(t *testing.T) (m map[string]feedTest) {
	files, err := filepath.Glob("../samples/*")
	if err != nil {
		t.Fatal(err)
	}

	var data []byte
	var format = "2006-01-02 15:04:05 -0700 MST"
	m = make(map[string]feedTest, len(files)/3)
	for _, file := range files {
		bits := strings.Split(filepath.Base(file), ".")
		name, ext := bits[0], bits[1]
		ft := m[name]
		switch ext {
		case "urls":
			if data, err = ioutil.ReadFile(file); err != nil {
				t.Fatal(err)
			}
			urls := bytes.Split(bytes.TrimRight(data, "\n"), []byte{'\n'})
			ft.URLs = make([]*url.URL, len(urls))
			for i := range urls {
				if ft.URLs[i], err = url.Parse(string(urls[i])); err != nil {
					t.Fatal(err)
				}
			}
		case "times":
			if data, err = ioutil.ReadFile(file); err != nil {
				t.Fatal(err)
			}
			times := bytes.Split(bytes.TrimRight(data, "\n"), []byte{'\n'})
			ft.Times = make([]time.Time, len(times))
			for i := range times {
				if ft.Times[i], err = time.Parse(format, string(times[i])); err != nil {
					t.Fatal(err)
				}
			}
		default:
			ft.Type = ext
			ft.Name = name
			ft.Path = file
			if ft.Data, err = ioutil.ReadFile(file); err != nil {
				t.Fatal(err)
			}
		}
		m[name] = ft
	}
	return
}
