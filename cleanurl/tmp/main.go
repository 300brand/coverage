package main

import (
	"encoding/json"
	"fmt"
	"github.com/300brand/logger"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

var AllURLs = make([]string, 0, 80000)

func init() {
	logger.Trace = log.New(ioutil.Discard, "", log.LstdFlags)
}

func main() {
	f, err := os.Open("/gocode/urls.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	if err := dec.Decode(&AllURLs); err != nil {
		log.Fatal(err)
	}

	// map[domain]map[badurl]goodurl
	tests := make(map[string]map[string]string)

	for _, rawurl := range AllURLs {
		dirty, err := url.Parse(rawurl)
		if err != nil {
			log.Fatal(err)
		}

		rawquery := dirty.RawQuery

		switch {
		case strings.Contains(rawquery, "fb="):
			rawquery = strings.Replace(rawquery, "fb=", "", 1)
		case strings.Contains(rawquery, "gplus="):
			rawquery = strings.Replace(rawquery, "gplus=", "", 1)
		case strings.Contains(rawquery, "CMP=OTC-RSS"):
			rawquery = strings.Replace(rawquery, "CMP=OTC-RSS", "", 1)
		case strings.Contains(rawquery, "&ref=25"):
			rawquery = strings.Replace(rawquery, "&ref=25", "", 1)
		case strings.Contains(rawquery, "ref=25"):
			rawquery = strings.Replace(rawquery, "ref=25", "", 1)
		case strings.Contains(rawquery, "cmp=RSS-FEED"):
			rawquery = strings.Replace(rawquery, "cmp=RSS-FEED", "", 1)
		case strings.Contains(rawquery, "asrc=RSS_BP"),
			strings.Contains(rawquery, "source=") && !strings.Contains(rawquery, "_source="):
			rawquery = ""
		case strings.Contains(rawquery, "ana=RSS&s=article_search"):
			rawquery = ""
		case strings.Contains(rawquery, "atc=") || strings.Contains(rawquery, "atc~c="):
			rawquery = ""
		case strings.Contains(rawquery, "op=1"):
			rawquery = strings.Replace(rawquery, "op=1", "", 1)
		case strings.Contains(rawquery, "cmpid="):
			rawquery = ""
		case strings.Contains(rawquery, "ncid=rss_semi"):
			rawquery = strings.Replace(rawquery, "ncid=rss_semi", "", 1)
		case strings.Contains(rawquery, "ex_cid"):
			rawquery = ""
		case strings.Contains(rawquery, "comm_ref=false&"):
			rawquery = strings.Replace(rawquery, "comm_ref=false&", "", 1)
		case strings.Contains(rawquery, "s_cid"):
			rawquery = ""
		case strings.Contains(rawquery, "lifehealth="):
			rawquery = ""
		case strings.Contains(rawquery, "feedName") && !strings.Contains(rawquery, "feedName=companyNews&feedType=RSS&"):
			rawquery = strings.Replace(rawquery, "amp=&", "", -1)
		case strings.Contains(rawquery, "feedName=companyNews&feedType=RSS&rpc=43&"):
			rawquery = strings.Replace(rawquery, "feedName=companyNews&feedType=RSS&rpc=43&", "", 1)
			rawquery = strings.Replace(rawquery, "amp=&", "", -1)
		case strings.Contains(rawquery, "feedName=companyNews&feedType=RSS&"):
			rawquery = strings.Replace(rawquery, "feedName=companyNews&feedType=RSS&", "", 1)
			rawquery = strings.Replace(rawquery, "amp=&", "", -1)
		case strings.Contains(rawquery, "mbid"):
			rawquery = ""
		case strings.Contains(rawquery, "f=102920358"):
			rawquery = ""
		case strings.Contains(rawquery, "f=1001"):
			rawquery = ""
		case strings.Contains(rawquery, "beta"):
			rawquery = ""
		case strings.Contains(rawquery, "cpage"):
			rawquery = strings.Replace(rawquery, "cpage=1&", "", 1)
		case strings.Contains(rawquery, "kc"):
			rawquery = ""
		case strings.Contains(rawquery, "urw="):
			rawquery = strings.Replace(rawquery, "urw=", "", 1)
		case strings.Contains(rawquery, "tc=page0"):
			rawquery = ""
		case strings.Contains(rawquery, ".99="):
			rawquery = ""
		case strings.Contains(rawquery, "dlvrit"):
			rawquery = ""
		case strings.Contains(rawquery, "attr=all"):
			rawquery = ""
		}

		dirty.RawQuery = rawquery

		if _, ok := tests[dirty.Host]; !ok {
			tests[dirty.Host] = make(map[string]string)
		}

		tests[dirty.Host][rawurl] = dirty.String()
	}

	letters := []byte(`1abcdefghilmnopqrstuvwy`)
	data := make(map[byte]io.Writer, len(letters))

	for _, ch := range letters {
		w, err := os.OpenFile(fmt.Sprintf("%s_test.go", []byte{ch}), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "package cleanurl\n\n")
		fmt.Fprintf(w, "import \"testing\"\n\n")
		defer w.Close()
		data[ch] = w
	}

	for domain, urls := range tests {

		unwww := strings.Replace(domain, "www.", "", 1)
		w, ok := data[unwww[0]]
		if !ok {
			log.Fatalf("No file for %s", domain)
		}

		fmt.Fprintf(w, "func Test_%s(t *testing.T) {\n", strings.Replace(strings.Replace(strings.Replace(domain, ".", "_", -1), "-", "_", -1), ":", "_", -1))
		fmt.Fprint(w, "\turls := map[string]string{\n")
		for bad, good := range urls {
			fmt.Fprintf(w, "\t\t"+`"%s": "%s",`+"\n", bad, good)
		}
		fmt.Fprint(w, "\t}\n\ttestURLs(t, urls)\n")
		fmt.Fprintf(w, "}\n")
	}
}
