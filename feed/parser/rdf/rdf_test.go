package rdf

import (
	"io/ioutil"
	"log"
	"testing"
	"time"
)

var bRDF, bAtom []byte

func init() {
	var err error
	if bRDF, err = ioutil.ReadFile("../../samples/NetworkWorld.rdf"); err != nil {
		log.Fatal(err)
	}
	if bAtom, err = ioutil.ReadFile("../../samples/TheRegister.atom"); err != nil {
		log.Fatal(err)
	}
}

func TestEntryLen(t *testing.T) {
	doc := Doc{}
	if err := doc.Decode(bRDF); err != nil {
		t.Error(err)
	}
	entries := doc.Item
	if len(entries) == 0 {
		t.Error("No entries found")
	}
	if len(entries) != 90 {
		t.Errorf("Invalid number of entries: %d", len(entries))
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
	if err := doc.Decode(bRDF); err != nil {
		t.Error(err)
	}
	if doc.Channel.Title == "" {
		t.Error("Blank title")
	}
	t.Logf("Title: %s", doc.Channel.Title)
}

func TestURLs(t *testing.T) {
	urls := []string{
		"http://www.networkworld.com/news/2012/111212-microsoft-to-address-yammer-integration-264178.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-verizon-droid-264180.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-microsoft-surface-rt-tablet-touch-264167.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-cray-bumps-ibm-from-top500-264168.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-windows-8-security-unshaken-by-264166.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-firefox-turns-8-and-gets-264165.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-nvidia-amd-release-graphics-processors-264169.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-rim-blackberry10-264160.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-pirate-bay-co-founder-now-suspected-264159.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-virnetx-targets-ipad-mini-iphone-264158.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-ebay-returns-to-china-with-264157.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-microsoft-likely-to-leash-ios-264156.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-skype-for-windows-phone-8-264155.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-patent-settlement-a-win-for-264153.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-intel-ships-60-core-xeon-phi-264152.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111212-un39s-civil-aviation-body-recommends-264151.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111112-apple-htc-settle-patent-suits-264150.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111012-google-access-returns-to-china-264149.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111012-space-station-commander-controls-lego-robot-264148.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/111012-manufacturer-sues-ibm-over-sap-264140.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-wall-street-beat-tech-mampa-264146.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-google-blocked-in-china-by-264144.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-three-simple-rules-for-buying-264143.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-staff-emails-are-not-owned-264142.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-cray-chases-clusters-with-appro-264138.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-nyc-corporate-partnership-seeks-fresh-264139.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-symantec-encryption-264137.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-crossbeam-thoma-bravo-264135.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-fatal-half-measures-in-incident-264128.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-iphone6-rumors-264127.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-adobe-reader-x-sandbox-bypassed-264129.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-ransomware-crooks-make-millions-from-264130.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-cisco-recommends-mcafee-switch-for-264121.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-china-could-be-behind-twitter-264123.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-shareholders-kept-in-the-dark-264122.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-judge-to-consider-samsung39s-questions-264124.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-iranian-minister-faces-us-sanctions-264118.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110912-michigan-man-charged-with-selling-264119.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-how-to-secure-big-data-264120.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-microsoft-launches-skype-centered-hub-for-264116.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-windows8-security-patch-264111.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-atampt-reverses-facetime-blocking-264117.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-csa-huawei-264110.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-android-264108.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-microsoft-slates-first-windows-8-264109.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-vmware-cloud-paas-264107.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-ipad-tips-and-tricks-for-264104.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-windows8-264106.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-smartphone-and-tablet-users-helped-264105.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-godaddy-apology-264100.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-iphone-foxconn-264099.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-hp-urges-consumer-customers-not-264102.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-isoc-facebook-comcast-ipv6-264096.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-twitter-asks-many-users-to-264097.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-tech-apologies-264095.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-samsung-galaxy-smartphone-264094.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-cisco-collaboration-264093.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-why-isnt-microsofts-answer-to-264092.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-the-cloud-as-data-center-264090.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-apple-university-264091.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-five-must-have-business-apps-for-264087.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-oracle-buys-instantis-for-project-264085.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-security-cloud-computing-264086.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-citrix-and-netapp-simplify-on-premises-264083.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-zero-day-pdf-exploit-reportedly-defeats-264084.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-social-nets-create-election39s-biggest-264082.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-hitachi-releases-new-16tb-flash-264081.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-verizon-expands-lync-service-to-264080.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-heist-once-again-highlights-e-banking-264077.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-siemens-industrial-software-targeted-by-264071.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-us-commission-fingers-china-as-264078.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-lenovo-expects-convertible-pcs-will-264079.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-joyent-264069.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-chinese-ex-hacker-says-working-for-264074.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-cray39s-next-supercomputer-has-speedy-264075.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-lenovo-sees-profit-growth-sag-264076.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-blackberry-10-is-fips-certified-264072.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110812-obama-tech-adviser-says-re-election-264073.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-design-supply-component-issues-may-264068.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-microsoft-office-ios-android-264067.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-oracle-hit-with-patent-lawsuit-264065.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-can-windows-8-give-developers-264064.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-it-salaries-2013-264063.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-amd-open-source-264062.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-twitter-a-big-winner-in-264059.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-aclu-eff-challenge-law-targeting-264060.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-us-election-could-mean-movement-264057.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-gartner-magic-quadrant-264058.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-samsung-laying-groundwork-for-server-264053.html?source=nww_rss",
		"http://www.networkworld.com/news/2012/110712-cloud-security-lawyers-264055.html?source=nww_rss",
	}
	doc := Doc{}
	if err := doc.Decode(bRDF); err != nil {
		t.Error(err)
	}
	entries := doc.Item
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
	// The feed provider really screwed the pooch here. TZ should be
	// America/New_York
	loc, err := time.LoadLocation("Canada/Atlantic")
	if err != nil {
		t.Error(err)
	}
	dates := []time.Time{
		time.Date(2012, time.November, 12, 12, 36, 0, 0, loc),
		time.Date(2012, time.November, 12, 12, 34, 25, 0, loc),
		time.Date(2012, time.November, 12, 10, 55, 0, 0, loc),
		time.Date(2012, time.November, 12, 10, 41, 0, 0, loc),
		time.Date(2012, time.November, 12, 10, 31, 0, 0, loc),
		time.Date(2012, time.November, 12, 10, 20, 0, 0, loc),
		time.Date(2012, time.November, 12, 10, 7, 0, 0, loc),
		time.Date(2012, time.November, 12, 9, 10, 18, 0, loc),
		time.Date(2012, time.November, 12, 8, 35, 50, 0, loc),
		time.Date(2012, time.November, 12, 7, 56, 3, 0, loc),
		time.Date(2012, time.November, 12, 7, 39, 57, 0, loc),
		time.Date(2012, time.November, 12, 7, 24, 0, 0, loc),
		time.Date(2012, time.November, 12, 7, 9, 31, 0, loc),
		time.Date(2012, time.November, 12, 5, 46, 13, 0, loc),
		time.Date(2012, time.November, 12, 1, 10, 5, 0, loc),
		time.Date(2012, time.November, 11, 10, 29, 42, 0, loc),
		time.Date(2012, time.November, 10, 9, 48, 5, 0, loc),
		time.Date(2012, time.November, 10, 12, 44, 4, 0, loc),
		time.Date(2012, time.November, 9, 9, 26, 5, 0, loc),
		time.Date(2012, time.November, 9, 8, 6, 28, 0, loc),
		time.Date(2012, time.November, 9, 6, 17, 25, 0, loc),
		time.Date(2012, time.November, 9, 5, 15, 0, 0, loc),
		time.Date(2012, time.November, 9, 3, 29, 0, 0, loc),
		time.Date(2012, time.November, 9, 3, 9, 0, 0, loc),
		time.Date(2012, time.November, 9, 3, 5, 53, 0, loc),
		time.Date(2012, time.November, 9, 2, 50, 0, 0, loc),
		time.Date(2012, time.November, 9, 2, 2, 38, 0, loc),
		time.Date(2012, time.November, 9, 1, 44, 47, 0, loc),
		time.Date(2012, time.November, 9, 11, 43, 0, 0, loc),
		time.Date(2012, time.November, 9, 11, 26, 10, 0, loc),
		time.Date(2012, time.November, 9, 10, 36, 0, 0, loc),
		time.Date(2012, time.November, 9, 10, 33, 0, 0, loc),
		time.Date(2012, time.November, 9, 10, 26, 32, 0, loc),
		time.Date(2012, time.November, 9, 7, 0, 0, 0, loc),
		time.Date(2012, time.November, 9, 7, 0, 0, 0, loc),
		time.Date(2012, time.November, 9, 6, 41, 0, 0, loc),
		time.Date(2012, time.November, 9, 3, 17, 0, 0, loc),
		time.Date(2012, time.November, 8, 11, 52, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 8, 0, 0, loc),
		time.Date(2012, time.November, 8, 6, 26, 0, 0, loc),
		time.Date(2012, time.November, 8, 5, 16, 10, 0, loc),
		time.Date(2012, time.November, 8, 5, 13, 0, 0, loc),
		time.Date(2012, time.November, 8, 5, 11, 57, 0, loc),
		time.Date(2012, time.November, 8, 4, 33, 45, 0, loc),
		time.Date(2012, time.November, 8, 4, 32, 0, 0, loc),
		time.Date(2012, time.November, 8, 4, 16, 13, 0, loc),
		time.Date(2012, time.November, 8, 4, 5, 0, 0, loc),
		time.Date(2012, time.November, 8, 3, 34, 1, 0, loc),
		time.Date(2012, time.November, 8, 3, 6, 0, 0, loc),
		time.Date(2012, time.November, 8, 2, 24, 18, 0, loc),
		time.Date(2012, time.November, 8, 1, 48, 5, 0, loc),
		time.Date(2012, time.November, 8, 1, 42, 0, 0, loc),
		time.Date(2012, time.November, 8, 12, 17, 15, 0, loc),
		time.Date(2012, time.November, 8, 12, 16, 0, 0, loc),
		time.Date(2012, time.November, 8, 11, 59, 20, 0, loc),
		time.Date(2012, time.November, 8, 11, 58, 20, 0, loc),
		time.Date(2012, time.November, 8, 11, 28, 32, 0, loc),
		time.Date(2012, time.November, 8, 10, 59, 0, 0, loc),
		time.Date(2012, time.November, 8, 10, 58, 0, 0, loc),
		time.Date(2012, time.November, 8, 10, 32, 53, 0, loc),
		time.Date(2012, time.November, 8, 10, 5, 0, 0, loc),
		time.Date(2012, time.November, 8, 9, 47, 0, 0, loc),
		time.Date(2012, time.November, 8, 9, 21, 9, 0, loc),
		time.Date(2012, time.November, 8, 8, 10, 0, 0, loc),
		time.Date(2012, time.November, 8, 8, 3, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 14, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 14, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 14, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 12, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 6, 0, 0, loc),
		time.Date(2012, time.November, 8, 7, 0, 0, 0, loc),
		time.Date(2012, time.November, 8, 6, 9, 0, 0, loc),
		time.Date(2012, time.November, 8, 6, 0, 0, 0, loc),
		time.Date(2012, time.November, 8, 5, 50, 0, 0, loc),
		time.Date(2012, time.November, 8, 4, 1, 0, 0, loc),
		time.Date(2012, time.November, 8, 3, 6, 0, 0, loc),
		time.Date(2012, time.November, 8, 1, 9, 0, 0, loc),
		time.Date(2012, time.November, 7, 9, 33, 0, 0, loc),
		time.Date(2012, time.November, 7, 5, 29, 0, 0, loc),
		time.Date(2012, time.November, 7, 4, 51, 41, 0, loc),
		time.Date(2012, time.November, 7, 4, 24, 0, 0, loc),
		time.Date(2012, time.November, 7, 4, 10, 0, 0, loc),
		time.Date(2012, time.November, 7, 3, 51, 42, 0, loc),
		time.Date(2012, time.November, 7, 3, 44, 9, 0, loc),
		time.Date(2012, time.November, 7, 3, 41, 0, 0, loc),
		time.Date(2012, time.November, 7, 3, 37, 0, 0, loc),
		time.Date(2012, time.November, 7, 3, 6, 0, 0, loc),
		time.Date(2012, time.November, 7, 2, 52, 12, 0, loc),
		time.Date(2012, time.November, 7, 1, 51, 0, 0, loc),
		time.Date(2012, time.November, 7, 1, 49, 37, 0, loc),
	}
	doc := Doc{}
	if err := doc.Decode(bRDF); err != nil {
		t.Error(err)
	}
	entries := doc.Item
	if len(entries) == 0 {
		t.Error("No entries found")
	}
	for i, e := range entries {
		if e.Date.Time().Unix() != dates[i].Unix() {
			t.Errorf("[%d] %s != %s", i, e.Date.Time(), dates[i])
		}
	}
}
