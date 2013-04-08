package bridge

import (
	"fmt"
	"git.300brand.com/coverage/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	s         *httptest.Server
	queueJSON = `{
	   "jsonrpc" : 2,
	   "error" : null,
	   "id" : 0,
	   "result" : [
	      {
	         "data" : "http://rss.techtarget.com/1532",
	         "object_id" : 14036,
	         "class" : "RemoveFeed",
	         "id" : 2751
	      },
	      {
	         "data" : "http://virtualization.sys-con.com/index.rss",
	         "object_id" : 14037,
	         "class" : "RemoveFeed",
	         "id" : 2752
	      },
	      {
	         "data" : "http://www.aacsb.edu/rss/bized.xml",
	         "object_id" : 14038,
	         "class" : "RemoveFeed",
	         "id" : 2753
	      },
	      {
	         "data" : "http://www.s-ox.com/dsp_getRSSMain.cfm",
	         "object_id" : 14039,
	         "class" : "RemoveFeed",
	         "id" : 2754
	      },
	      {
	         "data" : "https://twitter.com/statuses/user_timeline/27183959.rss",
	         "object_id" : 14040,
	         "class" : "RemoveFeed",
	         "id" : 2755
	      },
	      {
	         "data" : "http://feeds.feedburner.com/geekwire",
	         "object_id" : 14042,
	         "class" : "CoverageFeed",
	         "id" : 2756
	      },
	      {
	         "data" : {
	            "matches" : [],
	            "feeds" : [
	               "http://www.1to1media.com/weblog/atom.xml"
	            ],
	            "dateBounds" : {
	               "end" : "2013-01-29 17:14:45",
	               "start" : "2013-01-23 15:34:45"
	            },
	            "summaries" : [],
	            "id" : 14043,
	            "phrases" : {
	               "include" : [
	                  "CDW"
	               ],
	               "exclude" : []
	            },
	            "pages" : [
	               "http://www.slj.com/2013/01/industry-news/news-bites-apply-for-alsc-dia-mini-grants-by-february-1/",
	               "http://www.centerdigitaled.com/news/Tablets-Multi-Tasking.html",
	               "http://call-accounting.tmcnet.com/articles/317393-isis-telecom-audit-identifies-11m-savings-cdw-berbee.htm",
	               "http://www.virtual-strategy.com/2012/12/28/make-your-cloud-more-effective",
	               "http://redmondmag.com/articles/2012/10/30/organizations-underestimating-sharepoint-costs.aspx",
	               "http://redmondmag.com/articles/2012/07/01/nits-and-picks.aspx",
	               "http://washingtontechnology.com/articles/2013/01/23/sewp-v-draft-rfp-soon.aspx"
	            ]
	         },
	         "object_id" : 14043,
	         "class" : "CoverageReport",
	         "id" : 2757
	      },
	      {
	         "data" : "http://feeds.feedburner.com/chicagotribune/career",
	         "object_id" : 14059,
	         "class" : "RemoveFeed",
	         "id" : 2758
	      },
	      {
	         "data" : "http://us.generation-nt.com.feedsportal.com/c/33885/f/612606/index.rss",
	         "object_id" : 14060,
	         "class" : "CoverageFeed",
	         "id" : 2759
	      },
	      {
	         "data" : "http://www.ama-assn.org/amednews/newsfeed/rss20.rss",
	         "object_id" : 14061,
	         "class" : "CoverageFeed",
	         "id" : 2760
	      }
	   ]
	}`
)

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/queue", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, queueJSON)
	})
	s = httptest.NewServer(mux)
}

func TestQueue(t *testing.T) {
	config.RPC.Address = s.URL + "/queue"
	items, err := GetQueue(2750, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("items: %+v", items)
}
