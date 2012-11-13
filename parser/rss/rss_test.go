package rss

import (
	"git.300brand.com/coverage/parser"
	"git.300brand.com/coverage/parser/testfeed"
	"testing"
	"time"
)

func TestEntryLen(t *testing.T) {
	f := getRSSFeed(t)
	if len(f.Articles) != 10 {
		t.Errorf("Invalid number of entries: %d", len(f.Articles))
	}
}

func TestParseFail(t *testing.T) {
	rss := RSS{}
	if err := rss.Decode(testfeed.Atom); err == nil {
		t.Error("Expected error when parsing Atom feed")
	}
}

func TestTitle(t *testing.T) {
	f := getRSSFeed(t)
	if f.Title == "" {
		t.Error("Blank title")
	}
	t.Logf("Title: %s", f.Title)
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
	f := getRSSFeed(t)
	for i, a := range f.Articles {
		if a.URL.String() != urls[i] {
			t.Errorf("URL Mismatch:\nGOT: %s\nEXP: %s", a.URL.String(), urls[i])
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
	f := getRSSFeed(t)
	for i, a := range f.Articles {
		if !a.Published.Equal(dates[i]) {
			t.Errorf("[%d] %s != %s", i, a.Published, dates[i])
		}
	}

}

func getRSSFeed(t *testing.T) parser.Feed {
	rss := RSS{}
	err := rss.Decode(testfeed.RSS)
	if err != nil {
		t.Error(err)
	}
	return rss.Feed()
}
