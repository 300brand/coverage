package body

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

// http://www.biosciencetechnology.com/news/2013/03/brain-adds-specified-cells-during-puberty
// <div>Content</div><div>&nbsp;</div>
func TestBioscienceTechnology(t *testing.T) {
	//compareBodies(t, "BioscienceTechnology")
	t.Error("TODO")
}

// http://blogs.cio.com/security/17837/scientists-create-telepathic-rats-and-robosparrows
// Missing titles from article list at end of article - need to add \n after ul's li's
func TestCIOBlogs(t *testing.T) {
	//compareBodies(t, "CIOBlogs")
	t.Error("TODO")
}

// http://blog.cppionline.org/2012/10/supercharging-us-economy-where-both.html
// <div>Content</div><div><br></div>
func TestCPPIOnline(t *testing.T) {
	//compareBodies(t, "CPPIOnline")
	t.Error("TODO")
}

// http://www.medicalnewstoday.com/articles/257245.php
// Remove <select> tags
func TestMedicalNewsToday(t *testing.T) {
	//compareBodies(t, "MedicalNewsToday")
	t.Error("TODO")
}

// http://www.nytimes.com/2013/03/07/world/americas/after-chavezs-death-questions-about-future-abound-in-venezuela.html
// http://blog.scottlowe.org/2013/03/05/technology-short-take-30/
// http://mobileenterprise.edgl.com/how-to/The-Matrix-of-Mobile-Testing85037?rssid=Article85037
// http://www.tmcnet.com/viewette.aspx?u=http%3a%2f%2fwww.tmcnet.com%2fusubmit%2f2013%2f03%2f06%2f6971719.htm&kw=3
// http://www.pcmag.com/article2/0,2817,2416247,00.asp?kc=PCRSS03079TX1K0000585
// http://www.itbusiness.ca/it/client/en/Home/buildRss.asp?cid=67

// 221797 - Data.gov Expands Federal Communities, Global Impact
// http://gov.aol.com/2012/09/26/data-gov-expands-federal-communities-global-impact/
func TestAOLGov(t *testing.T) {
	compareBodies(t, "AOLGov")
}

// 340032 - Big Mac returns to SoCal as Dodgers hitting coach
// http://www.usatoday.com/story/sports/mlb/dodgers/2012/11/07/mark-mcgwire-hired-as-dodgers-hitting-coach/1690053/
// Updated to: http://www.usatoday.com/story/sports/mlb/dodgers/2012/11/07/mark-mcgwire-hired-as-dodgers-hitting-coach/1690053/?ajax=true
func TestUSAToday(t *testing.T) {
	compareBodies(t, "USAToday")
}

// 340038 - Greek Lawmakers Pass Austerity Deal
// http://online.wsj.com/article/SB10001424127887323894704578104832247833100.html
func TestWallStreetJournal(t *testing.T) {
	compareBodies(t, "WallStreetJournal")
}

// 340046 - Concur FY Q4 Revs Light; Profits Beat; Shrs Edge Higher
// http://www.forbes.com/sites/ericsavitz/2012/11/07/concur-fy-q4-revs-light-profits-beat-shrs-edge-higher/
func TestForbes(t *testing.T) {
	compareBodies(t, "Forbes")
}

// 340066 - Echo360 buys LectureTools in first acquisition since funding
// http://www.bizjournals.com/washington/news/2012/11/07/in-first-acquisition-since-funding.html?s=article_search
func TestBizJournals(t *testing.T) {
	//compareBodies(t, "BizJournals")
	t.Error("TODO")
}

// 340140 - Waze Launches In-App Advertising Platform
// http://www.pcmag.com/article2/0,2817,2411868,00.asp
func TestPCMag(t *testing.T) {
	//compareBodies(t, "PCMag")
	t.Error("TODO")
}

// 393887 - 2012 Federal Employee Viewpoint Survey Results
// http://gov.aol.com/2012/11/29/2012-federal-employee-viewpoint-survey-results/
func TestAOLGovSurveyResults(t *testing.T) {
	//compareBodies(t, "AOLGovSurveyResults")
	t.Error("TODO")
}

func compareBodies(t *testing.T, basename string) {
	// Open ze file!
	htmlIn, err := os.Open("../samples/" + basename + ".html")
	if err != nil {
		t.Error(err)
		return
	}
	defer htmlIn.Close()
	// Pull in ze data!
	html, err := ioutil.ReadAll(htmlIn)
	if err != nil {
		t.Error(err)
		return
	}
	// Clean ze HTMLz!
	clean, err := CleanHTML(html)
	if err != nil {
		t.Error(err)
		return
	}
	// Extract ze body!
	b, err := GetBody(clean)
	if err != nil {
		t.Error(err)
		return
	}
	// Open ze other file! The anticpiation, it grows!
	bodyIn, err := os.Open("../samples/" + basename + ".body")
	if err != nil {
		t.Error(err)
		return
	}
	defer bodyIn.Close()
	// Pull in ze other data!
	expect, err := ioutil.ReadAll(bodyIn)
	if err != nil {
		t.Error(err)
		return
	}
	// Trim to look pretty!
	expect = bytes.TrimSpace(expect)
	// Compare ze datas!
	if string(b.Text) != string(expect) {
		t.Error("Bodies do not match")
		t.Logf("Expect:\n%s", expect)
		t.Logf("Got:\n%s", b.Text)
	}
}
