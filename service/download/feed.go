package download

import (
	"git.300brand.com/coverage/downloader"
	//"git.300brand.com/coverage/logger"
	//"github.com/bketelsen/skynet"
	"errors"
	"time"
)

type FeedRequest struct {
	Timeout time.Duration
	Type    string
	URL     string
}

/*
func (s *DownloadService) Feed(ri *skynet.RequestInfo, req *FeedRequest, resp *parser.Feed) error {
	logger.Debug("DownloadService.Feed: ", req.URL)
}
*/

func downloadFeed(req *FeedRequest) (downloader.Response, error) {
	type pair struct {
		downloader.Response
		error
	}
	ch := make(chan pair, 1)
	go func() {
		r, err := downloader.Fetch(req.URL)
		ch <- pair{r, err}
	}()
	select {
	case <-time.After(req.Timeout):
		//logger.Warn("DownloadService.DownloadFeed: Timeout reached when downloading ", req.URL)
		return downloader.Response{}, errors.New("Timeout reached")
	case p := <-ch:
		return p.Response, p.error
	}
	panic("unreachable")
}
