package cleanurl

import (
	"net/url"
	"strings"
)

func Clean(u *url.URL) (out *url.URL) {
	out = u
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
		// Get rid of source=rss or track=rss
		case strings.Contains(k+strings.Join(v, ""), "rss"):
			values.Del(k)
		}
	}
	out.RawQuery = values.Encode()

	return
}
