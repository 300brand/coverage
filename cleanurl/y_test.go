package cleanurl

import "testing"

func Test_www_youtube_com(t *testing.T) {
	urls := map[string]string{
		"http://www.youtube.com/watch?feature=youtu.be&v=2NzYUzCOxZw":                                  "http://www.youtube.com/watch?feature=youtu.be&v=2NzYUzCOxZw",
		"http://www.youtube.com/watch?feature=youtu.be&v=QB_kgm2GZOU":                                  "http://www.youtube.com/watch?feature=youtu.be&v=QB_kgm2GZOU",
		"http://www.youtube.com/watch?feature=youtu.be&v=ZRJ67L3eriU":                                  "http://www.youtube.com/watch?feature=youtu.be&v=ZRJ67L3eriU",
		"http://www.youtube.com/watch?v=6rVdIu6Zmgw":                                                   "http://www.youtube.com/watch?v=6rVdIu6Zmgw",
		"http://www.youtube.com/watch?feature=c4-overview&list=UUBJycsmduvYEL83R_U4JriQ&v=CJTlB_S7ct0": "http://www.youtube.com/watch?feature=c4-overview&list=UUBJycsmduvYEL83R_U4JriQ&v=CJTlB_S7ct0",
		"http://www.youtube.com/watch?feature=plcp&v=UO-NhqFW9hM":                                      "http://www.youtube.com/watch?feature=plcp&v=UO-NhqFW9hM",
		"http://www.youtube.com/watch?feature=youtu.be&v=8zRqqAvIRTI":                                  "http://www.youtube.com/watch?feature=youtu.be&v=8zRqqAvIRTI",
		"http://www.youtube.com/watch?feature=youtu.be&v=xbNdZQ9smzY":                                  "http://www.youtube.com/watch?feature=youtu.be&v=xbNdZQ9smzY",
		"http://www.youtube.com/watch?v=S-RjR9oxQwE":                                                   "http://www.youtube.com/watch?v=S-RjR9oxQwE",
	}
	testURLs(t, urls)
}
