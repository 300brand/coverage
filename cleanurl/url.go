package cleanurl

import (
	"github.com/jbaikge/logger"
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
		// Remove all utm_ parameters
		case strings.HasPrefix(k, "utm_"):
			values.Del(k)
		// Remove rss= parameters
		case strings.Contains(k, "rss"):
			values.Del(k)
		// Get rid of source=rss or track=rss
		case filter(v, "rss"):
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
