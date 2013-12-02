package body

import (
	"github.com/moovweb/gokogiri"
	"testing"
)

func TestBrBr2P_Simple(t *testing.T) {
	data := []byte(`<div><br/><br/></div>`)
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		t.Fatal(err)
	}
	defer doc.Free()

	if err = BrBr2P(doc.Root()); err != nil {
		t.Fatal(err)
	}
	if got := doc.Root().ToUnformattedXml(); got != "<div/>" {
		t.Errorf("Got: %s", got)
	}
}

func TestBrBr2P_P(t *testing.T) {
	data := "<div><p>Monkey</p></div>"
	doc, err := gokogiri.ParseXml([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	defer doc.Free()

	if err = BrBr2P(doc.Root()); err != nil {
		t.Fatal(err)
	}
	if got := doc.Root().ToUnformattedXml(); got != string(data) {
		t.Errorf("Got: %s", got)
	}
}

func TestBrBr2P_Complex(t *testing.T) {
	data := []byte(`<div>
		Whitespace-prefixed text<br/><br/>
		Text with <a href="">a link</a><br/><br/>
		<!-- Random comment -->
		Single dumb break<br/>
		Text with <strong>Bold stuff</strong<br/><br/>
	</div>`)
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		t.Fatal(err)
	}
	defer doc.Free()

	if err = BrBr2P(doc.Root()); err != nil {
		t.Fatal(err)
	}
	expect := "<div><p>Whitespace-prefixed text</p>\n<p>Text with <a href=\"\">a link</a></p>\n<p>Single dumb break</p>\n<p>Text with <strong>Bold stuff</strong></p>\n</div>"
	if got := doc.Root().ToUnformattedXml(); got != expect {
		t.Errorf("Got: %s", got)
	}
}

func TestBrBr2P_Silly(t *testing.T) {
	data := []byte(`<div>
		one<br/>
		<br/>
		two<br/>
		<br/>
	</div>`)
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		t.Fatal(err)
	}
	defer doc.Free()

	if err = BrBr2P(doc.Root()); err != nil {
		t.Fatal(err)
	}
	expect := "<div><p>one</p>\n<p>two</p>\n</div>"
	if got := doc.Root().ToUnformattedXml(); got != expect {
		t.Errorf("Got: %s", got)
	}
}
