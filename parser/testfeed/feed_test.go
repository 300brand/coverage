package testfeed

import (
	"git.300brand.com/coverage/parser"
	_ "git.300brand.com/coverage/parser/atom"
	_ "git.300brand.com/coverage/parser/rdf"
	_ "git.300brand.com/coverage/parser/rss"
	"testing"
)

func TestAtomType(t *testing.T) {
	if typ, err := parser.Type(Atom); err != nil || typ != "atom" {
		t.Errorf("Expected atom, got %s", typ)
		t.Error(err)
	}
}

func TestRDFType(t *testing.T) {
	if typ, err := parser.Type(RDF); err != nil || typ != "rdf" {
		t.Errorf("Expected rdf, got %s", typ)
		t.Error(err)
	}
}

func TestRSSType(t *testing.T) {
	if typ, err := parser.Type(RSS); err != nil || typ != "rss" {
		t.Errorf("Expected rss, got %s", typ)
		t.Error(err)
	}
}

func TestInvalidType(t *testing.T) {
	data := []byte(`<?xml version="1.0"?><bunk><feed /></bunk>`)
	if typ, err := parser.Type(data); err == nil || typ != "" {
		t.Errorf("Expected blank type and error, got %s", typ)
		t.Error(err)
	}
}
