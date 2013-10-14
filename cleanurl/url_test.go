package cleanurl

import (
	"github.com/jbaikge/logger"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
	"testing"
)

func init() {
	logger.Trace = log.New(ioutil.Discard, "", log.LstdFlags)
}

func TestNoModify(t *testing.T) {
	base := []string{
		"http://asumag.com/Construction/planning/EIS-jury-commentary-201208/index.html?imw=Y",
		"http://blogs.wsj.com/washwire/2012/09/19/political-wisdom-fallout-from-romney-video/?mod=WSJBlog",
		"http://content.usatoday.com/communities/gameon/post/2012/08/dwight-howard-lakers-introductory-press-conference/1?csp=34sports",
		"http://oukc.oracle.com/static05/opn/login/?c=1188150698&t=offering",
		"http://pharmaceuticalcommerce.com/special_report?articleid=1731",
		"http://thejournal.com/articles/2008/08/26/school-security--august-26-2008.aspx",
		"http://www.aasa.org/content.aspx?id=12710",
		"http://www.destinationcrm.com/Articles/ReadArticle.aspx?ArticleID=73449",
		"http://www.economist.com/node/21559734?fsrc=gn_ep",
		"http://www.forbes.com/sites/energystockchannel/2012/09/17/questar-str-shares-cross-below-200-dma/",
		"http://www.hhnmag.com/hhnmag/HHNDaily/HHNDailyDisplay.dhtml?id=9650009405",
		"http://www.hreonline.com/HRE/story.jsp?storyId=533350286&topic=Main",
		"http://www.informationweek.com/byte/quickview/1485?wc=4",
		"http://www.oracle.com/us/corporate/press/1672622?ssSourceSiteId=opn",
		"http://www.ttnews.com/articles/basetemplate.aspx?storyid=30204",
		"http://www.usatoday.com/news/education/story/2012-08-08/dc-schools-cheating/56888224/1?csp=34news",
	}
	urls := make(map[string]string, len(base))
	for _, u := range base {
		urls[u] = u
	}
	testURLs(t, urls)
}

func TestNum(t *testing.T) {
	urls := map[string]string{
		"http://my.tld?123":     "http://my.tld?123",
		"http://my.tld?123&456": "http://my.tld?123&456",
	}
	testURLs(t, urls)
}

func TestRSS(t *testing.T) {
	urls := map[string]string{
		"http://my.tld?rss=123":           "http://my.tld",
		"http://my.tld?param=rss":         "http://my.tld",
		"http://my.tld?param=1&param=2":   "http://my.tld?param=1&param=2",
		"http://my.tld?param=rss&param=2": "http://my.tld",
	}
	testURLs(t, urls)
}

func TestUTM(t *testing.T) {
	urls := map[string]string{
		"http://technology.amis.nl/2012/08/08/quick-introduction-of-what-and-why-of-oracle-database-edition-based-redefinition/?utm_source=rss&utm_medium=rss&utm_campaign=quick-introduction-of-what-and-why-of-oracle-database-edition-based-redefinition": "http://technology.amis.nl/2012/08/08/quick-introduction-of-what-and-why-of-oracle-database-edition-based-redefinition/",
		"http://www.extremetech.com/electronics/135327-ifa-move-over-3d-its-time-for-4k-uhdtv?utm_source=rss&utm_medium=rss&utm_campaign=ifa-move-over-3d-its-time-for-4k-uhdtv":                                                                             "http://www.extremetech.com/electronics/135327-ifa-move-over-3d-its-time-for-4k-uhdtv",
		"http://www.fiercehealthcare.com/story/nurses-catch-more-medication-mistakes-supportive-hospitals/2012-08-30?utm_source=rss&utm_medium=rss":                                                                                                          "http://www.fiercehealthcare.com/story/nurses-catch-more-medication-mistakes-supportive-hospitals/2012-08-30",
		"http://www.fiercehealthit.com/story/thefts-stanford-oregon-hospitals-jeopardize-patient-info-nearly-17k/2012-08-06?utm_source=rss":                                                                                                                  "http://www.fiercehealthit.com/story/thefts-stanford-oregon-hospitals-jeopardize-patient-info-nearly-17k/2012-08-06",
		"http://www.modernhealthcare.com/article/20120913/NEWS/120919990?utm_source=rss01&utm_medium=rss&utm_campaign=rss01":                                                                                                                                 "http://www.modernhealthcare.com/article/20120913/NEWS/120919990",
	}
	testURLs(t, urls)
}

func TestWashingtonPost(t *testing.T) {
	urls := map[string]string{
		"http://www.washingtonpost.com/blogs/plum-line/post/budget-war-showcases-sharp-contrast-in-values-and-priorities/2013/03/12/03c9de08-8b3c-11e2-9b1a-deb258a24f2d_blog.html?wprss=rss_opinions": "http://www.washingtonpost.com/blogs/plum-line/post/budget-war-showcases-sharp-contrast-in-values-and-priorities/2013/03/12/03c9de08-8b3c-11e2-9b1a-deb258a24f2d_blog.html",
	}
	testURLs(t, urls)
}

func testURLs(t *testing.T, urls map[string]string) {
	for u, exp := range urls {
		test, err := url.Parse(u)
		if err != nil {
			t.Errorf("[%s] %s", u, err)
		}
		expect, err := url.Parse(exp)
		if err != nil {
			t.Errorf("[%s] %s", exp, err)
		}
		cleaned := Clean(test)
		// Quick equality check
		if expect.String() == cleaned.String() {
			continue
		}
		// Quick query length check
		if len(expect.Query()) != len(cleaned.Query()) {
			t.Errorf("[%s] invalid number of query parameters - Expect: %d; Cleaned: %d", test, len(expect.Query()), len(cleaned.Query()))
			continue
		}
		t.Log("Expect:  ", expect.RawQuery)
		t.Log("Cleaned: ", cleaned.RawQuery)
		// Deep query equality check
		for k, v := range expect.Query() {
			c, ok := cleaned.Query()[k]
			if !ok {
				t.Errorf("[%s] Couldn't find key '%s' in cleaned", test, k)
				continue
			}
			vj, cj := strings.Join(v, "&"), strings.Join(c, "&")
			if vj != cj {
				t.Errorf("[%s] %s != %s", test, vj, cj)
			}
		}
	}
}
