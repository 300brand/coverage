package rss

import (
	"io/ioutil"
	"log"
	"testing"
	"time"
)

var bRSS, bAtom []byte

func init() {
	var err error
	if bRSS, err = ioutil.ReadFile("../../samples/NASA.rss"); err != nil {
		log.Fatal(err)
	}
	if bAtom, err = ioutil.ReadFile("../../samples/TheRegister.atom"); err != nil {
		log.Fatal(err)
	}
}

func TestEntryLen(t *testing.T) {
	doc := Doc{}
	if err := doc.Decode(bRSS); err != nil {
		t.Error(err)
	}
	if len(doc.Channel.Item) != 10 {
		t.Errorf("Invalid number of entries: %d", len(doc.Channel.Item))
	}
}

func TestParseFail(t *testing.T) {
	doc := Doc{}
	if err := doc.Decode(bAtom); err == nil {
		t.Error("Expected error when parsing Atom feed")
	}
}

func TestTitle(t *testing.T) {
	doc := Doc{}
	if err := doc.Decode(bRSS); err != nil {
		t.Error(err)
	}
	if doc.Channel.Title == "" {
		t.Error("Blank title")
	}
	t.Logf("Title: %s", doc.Channel.Title)
}

func TestURLs(t *testing.T) {
	urls := []string{
		"http://www.nasa.gov/home/hqnews/2012/oct/HQ_12-387_Mars_Atmosphere.html",
		"http://www.nasa.gov/home/hqnews/2012/oct/HQ_12-384_Spot_Station.html",
		"http://www.nasa.gov/home/hqnews/2012/nov/HQ_C12-056_Ames_Safety_Contract.html",
		"http://www.nasa.gov/home/hqnews/2012/nov/HQ_12-386_SLS_RFI.html",
		"http://www.nasa.gov/home/hqnews/2012/nov/HQ_12-385_Fermi_Fog.html",
		"http://www.nasa.gov/home/hqnews/2012/nov/HQ_12-378_SpaceX_Commercial_Crew_Milestones.html",
		"http://www.nasa.gov/home/hqnews/2012/oct/HQ_C12-058_Ames_AE_Svcs.html",
		"http://www.nasa.gov/home/hqnews/2012/oct/HQ_C12-057_GSFC_NSSC_SAIC_Mod.html",
		"http://www.nasa.gov/home/hqnews/2012/oct/HQ_M12-211_Mars_Atmosphere_Telecon.html",
		"http://www.nasa.gov/home/hqnews/2012/oct/HQ_12-377_NASA-WPI_2013_Robot_Competition_Registration.html",
	}
	doc := Doc{}
	if err := doc.Decode(bRSS); err != nil {
		t.Error(err)
	}
	entries := doc.Channel.Item
	if len(entries) == 0 {
		t.Error("No entries found")
	}
	for i, e := range entries {
		if e.Link != urls[i] {
			t.Errorf("URL Mismatch:\nGOT: %s\nEXP: %s", e.Link, urls[i])
		}
	}
}

func TestTimestamps(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error(err)
	}
	dates := []time.Time{
		time.Date(2012, time.November, 2, 0, 0, 0, 0, loc),
		time.Date(2012, time.November, 2, 0, 0, 0, 0, loc),
		time.Date(2012, time.November, 1, 0, 0, 0, 0, loc),
		time.Date(2012, time.November, 1, 0, 0, 0, 0, loc),
		time.Date(2012, time.November, 1, 0, 0, 0, 0, loc),
		time.Date(2012, time.November, 1, 0, 0, 0, 0, loc),
		time.Date(2012, time.October, 31, 0, 0, 0, 0, loc),
		time.Date(2012, time.October, 31, 0, 0, 0, 0, loc),
		time.Date(2012, time.October, 31, 0, 0, 0, 0, loc),
		time.Date(2012, time.October, 31, 0, 0, 0, 0, loc),
	}
	doc := Doc{}
	if err := doc.Decode(bRSS); err != nil {
		t.Error(err)
	}
	entries := doc.Channel.Item
	if len(entries) == 0 {
		t.Error("No entries found")
	}
	for i, e := range entries {
		if !e.PubDate.Time().Equal(dates[i]) {
			t.Errorf("[%d] %s != %s", i, e.PubDate, dates[i])
		}
	}

}
