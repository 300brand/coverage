package validurl

import (
	"github.com/300brand/logger"
	"io/ioutil"
	"log"
	"net/url"
	"testing"
)

func init() {
	logger.Trace = log.New(ioutil.Discard, "", log.LstdFlags)
}

func TestInvalidURLs(t *testing.T) {
	urls := map[string]bool{
		"http://www.auntminnie.com/index.aspx?sec=def":  false,
		"http://cdn.auntminnie.com/rss/rss.aspx":        false,
		"https://accounts.google.com":                   false,
		"https://login.oracle.com/mysso/signon.jsp":     false,
		"http://apex.oracle.com/i/index.html":           false,
		"http://www.aviationweek.com/":                  false,
		"http://aviationweek.com":                       false,
		"http://www.businesswire.com/portal/site/home/": false,
		"http://businesswire.com/portal/site/home/":     false,
		"http://blogs.ft.com":                           false,
		"http://www.ft.com/comment/blogs":               false,
		"http://ft.com/comment/blogs":                   false,
		"http://www.complianceweek.com":                 false,
		"http://complianceweek.com":                     false,
		"https://app.cvent.com/Subscribers/Login.aspx":  false,
		"http://www.cvent.com":                          false,
		"http://cvent.com":                              false,
		"http://ftalphaville.ft.com/tag/alphachat/":     false,
		"http://www.ft.com":                             false,
		"http://ft.com":                                 false,
		"https://www.gartner.com/login/loginInitAction.do?method=initialize":                 false,
		"http://gartner.com/login/loginInitAction.do?method=initialize":                      false,
		"http://globalgoodgroup.com/feed/":                                                   false,
		"https://spin.infoedglobal.com/Home/Authorize":                                       false,
		"http://www.infoed.org/GeniusSearch/genius.asp":                                      false,
		"https://items.infoedglobal.com/Conversions/clientlogin.asp":                         false,
		"https://www.yammer.com/icagforum/":                                                  false,
		"http://infoedglobal.com":                                                            false,
		"http://www.insidehealth.com/feed/":                                                  false,
		"http://insidehealth.com/feed/":                                                      false,
		"http://investor.techtarget.com/CorporateProfile.aspx?iid=4151732":                   false,
		"http://www.lohud.com":                                                               false,
		"http://lohud.com":                                                                   false,
		"http://www.medtechwebinars.com":                                                     false,
		"http://medtechwebinars.com":                                                         false,
		"http://www.mightydeals.com":                                                         false,
		"http://mightydeals.com":                                                             false,
		"http://news.cincinnati.com/needlogin":                                               false,
		"https://oraclemeetings.webex.com/mw0307l/mywebex/default.do?siteurl=oraclemeetings": false,
		"http://oukc.oracle.com":                                                             false,
		"http://education.oracle.com/pls/web_prod-plq-dad/db_pages.getpage?page_id=160":      false,
		"https://oraclecse.webex.com/mw0307l/mywebex/default.do?siteurl=oraclecse":           false,
		"https://oracleus.activeevents.com/oracle_maintenance.html":                          false,
		"https://oracletalk.webex.com/mw0307l/mywebex/default.do?siteurl=oracletalk":         false,
		"http://oracle.6connex.com":                                                          false,
		"http://reg.accelacomm.com":                                                          false,
		"http://solutioncenters.computerworld.com":                                           false,
		"http://search.itunes.apple.com":                                                     false,
		"http://searchcloudprovider.techtarget.com":                                          false,
		"http://www.shrm.org/Pages/default.aspx":                                             false,
		"http://shrm.org/Pages/default.aspx":                                                 false,
		"http://searchconsumerization.techtarget.com":                                        false,
		"http://searchitchannel.techtarget.com":                                              false,
		"http://us.rd.yahoo.com":                                                             false,
		"http://yahoo.com":                                                                   false,
		"http://www.yahoo.com":                                                               false,
		"http://windowsitpro.com/private-cloud/not-surprising-businesses-focus-private-cloud-public-cloud-not-viable":     false,
		"http://www.windowsitpro.com/private-cloud/not-surprising-businesses-focus-private-cloud-public-cloud-not-viable": false,
		"http://www-nc.nytimes.com":                                                                       false,
		"https://myaccount.nytimes.com/auth/login?URI=www-nc.nytimes.com/&REFUSE_COOKIE_ERROR=SHOW_ERROR": false,
		"http://www3.ambest.com/ambv/bestnews/newscontent.aspx?refnum=169823&AltSrc=23":                   false,
		"http://www.ambest.com/":                                                                          false,
		"http://ambest.com/":                                                                              false,
	}
	testURLs(t, urls)
}

func TestValidURLs(t *testing.T) {
	urls := map[string]bool{
		"http://feeds.feedburner.com/yankodesign":                                                                 true,
		"http://blogs.windows.com/windows/b/bloggingwindows/rss.aspx":                                             true,
		"http://blog.twitter.com/feeds/posts/default?alt=rss":                                                     true,
		"http://rss.tmcnet.com/rss/crss.ashx?cnl=Master+Agent":                                                    true,
		"http://feeds2.feedburner.com/the_tech_labs":                                                              true,
		"http://www.techinvestornews.com/Facebook/Latest-Facebook-News-RSS":                                       true,
		"http://mobilesyrup.com/feed/":                                                                            true,
		"http://feeds.mercurynews.com/mngi/rss/CustomRssServlet/568/200763.xml":                                   true,
		"http://www.miamiherald.com/news/columnists/fred-grimm/index.xml":                                         true,
		"http://www.dispatch.com/content/blogs/a-look-back/index.html?rss":                                        true,
		"http://rss.cincinnati.com/apps/pbcs.dll/section?category=rssenq02&mime=xml":                              true,
		"http://rss.techtarget.com/444.xml":                                                                       true,
		"http://feeds.chicagotribune.com/chicagotribune/cars/":                                                    true,
		"http://www.lifescienceleader.com/component/fpss/module/156?format=feed&type=rss":                         true,
		"http://biosciencetechnology.com/rss.aspx?id=91":                                                          true,
		"http://www.americanbanker.com/resources/mortgages.xml":                                                   true,
		"http://www.powermag.com/rss/renewables/wind/rss.xml":                                                     true,
		"http://rcpmag.com/rss-feeds/news.aspx":                                                                   true,
		"http://www.windowsitpro.com/topviewed/default.aspx?rss=TopViewed":                                        true,
		"http://washingtontimes.feedsportal.com/c/34503/fe.ed/www.washingtontimes.com/rss/weblogs/wizards-watch/": true,
	}
	testURLs(t, urls)
}

func testURLs(t *testing.T, urls map[string]bool) {
	for tu, exp := range urls {
		u, err := url.Parse(tu)
		if err != nil {
			log.Fatal(err)
		}
		if IsValid(u) == exp {
			t.Log("Passed: ", u)
		} else {
			t.Log("Failed: ", u)
		}
	}
}
