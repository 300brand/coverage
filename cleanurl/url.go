package cleanurl

import (
	"github.com/300brand/logger"
	"net/url"
	"strings"
)

func Clean(u *url.URL) (out *url.URL) {
	out = new(url.URL)

	defer logger.Trace.Printf("Clean: result: %s", out)
	logger.Trace.Printf("Clean: called %s", u)

	*out = *u

	out.Fragment = ""

	// For The VAR Guy - Drop down to http and normalize host
	if strings.HasSuffix(u.Host, "thevarguy.com") {
		out.Scheme, out.Host = "http", "thevarguy.com"
	}
	// Fix computeruser query string
	if u.Host == "www.computeruser.com" {
		out.RawQuery = strings.Split(u.RawQuery, "%3F")[0]
		return
	}
	// Solve nytimes redirect to login
	if u.Host == "www-nc.nytimes.com" {
		out.Host = "www.nytimes.com"
	}
	// Fix reuters host
	if strings.HasSuffix(u.Host, "reuters.com") {
		out.Host = "reuters.com"
	}
	// Fix infoworld host
	if strings.HasSuffix(u.Host, "infoworld.com") {
		out.Host = "infoworld.com"
	}
	values := u.Query()

	// bounce if there's nothing to do
	if len(values) == 0 {
		return
	}

	for k, v := range values {

		switch {
		// Remove all utm_ / utm- parameters
		case strings.HasPrefix(k, "utm"),
			// Remove rss= parameters
			strings.Contains(k, "rss"),
			// Get rid of source=rss or track=rss
			filter(v, "rss"),
			// Remove fb= parameters
			strings.HasPrefix(k, "fb"),
			// news.cnet.com adds some weird stuff
			k == "tag",
			k == "subj",
			// Remove gplus= parameters
			k == "gplus",
			// Remove CMP= and cmp= parameters
			k == "CMP",
			k == "cmp",
			// Get rid of ref=25
			k == "ref" && filter(v, "25"),
			// Remove asrc= parameters
			k == "asrc",
			// Remove source= parameters
			k == "source",
			// Get rid of ana=RSS
			k == "ana" && filter(v, "RSS"),
			// Get rid of s=article_search
			k == "s" && filter(v, "article_search"),
			// Remove atc~c= and atc= parameters
			strings.HasPrefix(k, "atc"),
			// Remove op= parameters
			k == "op",
			// Remove cmpid= parameters
			k == "cmpid",
			// Remove sc= parameters
			k == "sc",
			// Remove ncid= parameters
			k == "ncid",
			// Remove ex_cid= parameters
			k == "ex_cid",
			// Remove comm_ref= parameters
			k == "comm_ref",
			//Remove s_cid= parameters
			k == "s_cid",
			// Remove lifehealth= parameters
			k == "lifehealth",
			// Remove feedName= parameters
			k == "feedName",
			// Remove feedType= parameters
			k == "feedType",
			// Remove virtualBrandChannel= parameters
			k == "virtualBrandChannel",
			// Remove mbid= parameters
			k == "mbid",
			// Remove f= parameters
			k == "f",
			// Remove ft= parameters
			k == "ft",
			// Remove beta= parameters
			k == "beta",
			// Remove action_object_map= parameters
			k == "action_object_map",
			// Remove action_ref_map= parameters
			k == "action_ref_map",
			// Remove action_type_map= parameters
			k == "action_type_map",
			// Remove cpage= parameters
			k == "cpage",
			// Remove kc= parameters
			k == "kc",
			// Remove urw= parameters
			k == "urw",
			// Remove tc= parameters
			k == "tc",
			// Remove %s.99= parameters
			strings.Contains(k, ".99"),
			// Remove dlvrit= parameters
			k == "dlvrit",
			// Get rid of attr=all
			k == "attr" && filter(v, "all"),
			// Remove logvisit= parameters
			k == "logvisit",
			// Remove npu= parameters
			k == "npu",
			// Remove rpc= parameters
			k == "rpc",
			// Remove amp= parameters
			k == "amp",
			// Remove = parameters
			k == "",
			// Remove _r= parameters
			k == "_r":
			// Get rid of type=companyNews
			// k == "type" && filter(v, "companyNews"):
			values.Del(k)
		}
	}
	out.RawQuery = values.Encode()

	return
}

func filter(haystack []string, needle string) (found bool) {
	for _, v := range haystack {
		found = found || v == needle
	}
	return
}
