package parser

import (
	"git.300brand.com/coverage/parser/testfeed"
	"testing"
)

func TestRSSEntryLen(t *testing.T) {
	f := getRSSFeed(t)
	if len(f.Articles) != 10 {
		t.Errorf("Invalid number of entries: %d", len(f.Articles))
	}
}

func TestRSSInit(t *testing.T) {
	if _, ok := decoders["Atom"]; !ok {
		t.Error("Atom decoder not found")
	}
}

func TestRSSFail(t *testing.T) {
	a := decoders["RSS"]
	if _, err := a.Decode(testfeed.Atom); err == nil {
		t.Error("Expected error when parsing Atom feed")
	}
}

func TestRSSTitle(t *testing.T) {
	f := getRSSFeed(t)
	if f.Title == "" {
		t.Error("Blank title")
	}
	t.Logf("Title: %s", f.Title)
}

func TestRSSURLs(t *testing.T) {
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
	f := getRSSFeed(t)
	for i, e := range f.Articles {
		if e.URL.String() != urls[i] {
			t.Errorf("URL Mismatch:\nGOT: %s\nEXP: %s", e.URL.String(), urls[i])
		}
	}
}

func getRSSFeed(t *testing.T) Feed {
	a := decoders["RSS"]
	f, err := a.Decode(testfeed.RSS)
	if err != nil {
		t.Error(err)
	}
	return f
}
