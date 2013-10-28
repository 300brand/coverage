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
	values := u.Query()

	// bounce if there's nothing to do
	if len(values) == 0 {
		return
	}

	for k, v := range values {
		switch {
		// Remove all utm_ / utm- parameters
		case strings.Contains(k, "utm_"),
			strings.Contains(k, "utm-"),
			// Remove rss= parameters
			strings.Contains(k, "rss"),
			// Get rid of source=rss or track=rss
			filter(v, "rss"),
			// Remove fb= parameters
			strings.Contains(k, "fb"),
			// Remove gplus= parameters
			strings.Contains(k, "gplus"),
			// Remove CMP= and cmp= parameters
			strings.Contains(k, "CMP"),
			strings.Contains(k, "cmp"),
			// Get rid of ref=25
			strings.Contains(k, "ref") && filter(v, "25"),
			// Remove asrc= parameters
			strings.Contains(k, "asrc"),
			// Remove source= parameters
			k == "source",
			// Get rid of ana=RSS
			strings.Contains(k, "ana") && filter(v, "RSS"),
			// Get rid of s=article_search
			strings.Contains(k, "s") && filter(v, "article_search"),
			// Remove atc~c= and atc= parameters
			strings.Contains(k, "atc"),
			// Remove op= parameters
			k == "op",
			// Remove cmpid= parameters
			strings.Contains(k, "cmpid"),
			// Remove sc= parameters
			strings.Contains(k, "sc"),
			// Remove ncid= parameters
			strings.Contains(k, "ncid"),
			// Remove ex_cid= parameters
			strings.Contains(k, "ex_cid"),
			// Remove comm_ref= parameters
			strings.Contains(k, "comm_ref"),
			//Remove s_cid= parameters
			strings.Contains(k, "s_cid"),
			// Remove lifehealth= parameters
			strings.Contains(k, "lifehealth"),
			// Remove feedName= parameters
			strings.Contains(k, "feedName"),
			// Remove feedType= parameters
			strings.Contains(k, "feedType"),
			// Remove virtualBrandChannel= parameters
			strings.Contains(k, "virtualBrandChannel"),
			// Remove mbid= parameters
			strings.Contains(k, "mbid"),
			// Get rid of f=102920358
			strings.Contains(k, "f") && filter(v, "102920358"),
			// Get rid of f=1001
			strings.Contains(k, "f") && filter(v, "1001"),
			// Get rid of ft=1
			strings.Contains(k, "ft") && filter(v, "1"),
			// Remove beta= parameters
			strings.Contains(k, "beta"),
			// Remove action_object_map= parameters
			strings.Contains(k, "action_object_map"),
			// Remove action_ref_map= parameters
			strings.Contains(k, "action_ref_map"),
			// Remove action_type_map= parameters
			strings.Contains(k, "action_type_map"),
			// Remove cpage= parameters
			strings.Contains(k, "cpage"),
			// Remove kc= parameters
			strings.Contains(k, "kc"),
			// Remove urw= parameters
			strings.Contains(k, "urw"),
			// Remove tc= parameters
			k == "tc",
			// Remove %s.99= parameters
			strings.Contains(k, ".99"),
			// Remove dlvrit= parameters
			strings.Contains(k, "dlvrit"),
			// Get rid of attr=all
			strings.Contains(k, "attr") && filter(v, "all"),
			// Remove logvisit= parameters
			strings.Contains(k, "logvisit"),
			// Remove npu= parameters
			strings.Contains(k, "npu"),
			// Get rid of rpc=43
			strings.Contains(k, "rpc") && filter(v, "43"),
			// Remove amp= parameters
			k == "amp":
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
