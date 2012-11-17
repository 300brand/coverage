package testfeed

import (
	"git.300brand.com/coverage/parser"
	_ "git.300brand.com/coverage/parser/atom"
	"git.300brand.com/coverage/parser/normalizer"
	_ "git.300brand.com/coverage/parser/rdf"
	_ "git.300brand.com/coverage/parser/rss"
	"testing"
)

var tests = []struct {
	Data []byte
	Type string
	Len  int
}{
	{
		Data: Atom,
		Type: "atom",
		Len:  50,
	},
	{
		Data: RDF,
		Type: "rdf",
		Len:  90,
	},
	{
		Data: RSS,
		Type: "rss",
		Len:  10,
	},
}

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
