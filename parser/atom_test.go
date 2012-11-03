package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

var atomFeed []byte

func init() {
	t := &testing.T{}
	f, err := os.Open("headlines.atom")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	atomFeed, err = ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
}

func TestEntries(t *testing.T) {
	f := getFeed(t)
	t.Log(len(f.Articles))
}

func TestInit(t *testing.T) {
	if _, ok := decoders["Atom"]; !ok {
		t.Error("Atom decoder not found")
	}
}

func TestTitle(t *testing.T) {
	f := getFeed(t)
	if f.Title == "" {
		t.Error("Blank title")
	}
	t.Logf("Title: %s", f.Title)
}

func getFeed(t *testing.T) Feed {
	a := decoders["Atom"]
	f, err := a.Decode(atomFeed)
	if err != nil {
		t.Error(err)
	}
	return f
}
