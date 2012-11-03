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

func TestInit(t *testing.T) {
	if _, ok := decoders["Atom"]; !ok {
		t.Error("Atom decoder not found")
	}
}

func TestTitle(t *testing.T) {
	a := decoders["Atom"]
	f, err := a.Decode(atomFeed)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Title: %s", f.Title)
}
