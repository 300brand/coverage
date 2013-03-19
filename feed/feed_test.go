package feed

import (
	//"git.300brand.com/coverage/feed/normalizer"
	//"git.300brand.com/coverage/feed/parser"
	"testing"

	_ "git.300brand.com/coverage/feed/parser/atom"
	_ "git.300brand.com/coverage/feed/parser/rdf"
	_ "git.300brand.com/coverage/feed/parser/rss"
	"log"
	"path/filepath"
)

type test struct {
	Files struct {
		In   []string
		URLs []string
	}
	Type string
	Len  int
}

var tests = []test{}

func init() {
	for _, ext := range []string{"atom", "rdf", "rss"} {
		list, err := filepath.Glob("samples/*." + ext)
		if err != nil {
			log.Fatal(err)
		}
		for _, filename := range list {
			log.Println(filename)
		}
	}
}

func TestURLs(t *testing.T) {}

/*
func TestParse(t *testing.T) {
	for _, test := range tests {
		n := &normalizer.Default{}
		if err := parser.Normalize(test.Data, n); err != nil {
			t.Error(err)
			continue
		}
		if len(n.Articles) != test.Len {
			t.Errorf("Invalid number of articles: %d", len(n.Articles))
		}
	}
}

func TestType(t *testing.T) {
	for _, test := range tests {
		if typ, err := parser.Type(test.Data); err != nil || typ != test.Type {
			t.Errorf("Expected %s, got %s", test.Type, typ)
			t.Error(err)
		}
	}
}

func TestInvalidParse(t *testing.T) {
	if _, err := parser.ParseType(Atom, "rss"); err == nil {
		t.Error("Expected error when using RSS decoder to parse Atom feed")
	}
}

func TestInvalidType(t *testing.T) {
	data := []byte(`<?xml version="1.0"?><bunk><feed /></bunk>`)
	if typ, err := parser.Type(data); err == nil || typ != "" {
		t.Errorf("Expected blank type and error, got %s", typ)
		t.Error(err)
	}
}
*/
