package download

import (
	"git.300brand.com/coverage/downloader"
	"git.300brand.com/coverage/logger"
	//"git.300brand.com/coverage/parser"
	//"github.com/bketelsen/skynet"
	"errors"
	"time"
)

type FeedRequest struct {
	URL     string
	Timeout int64
}

/*
func (s *DownloadService) Feed(ri *skynet.RequestInfo, req *FeedRequest, resp *parser.Feed) error {
	logger.Debug("DownloadService.Feed: ", req.URL)
}
*/

func downloadFeed(url string, timeout time.Duration) (downloader.Response, error) {
	type pair struct {
		downloader.Response
		error
	}
	ch := make(chan pair, 1)
	go func() {
		r, err := downloader.Fetch(url)
		ch <- pair{r, err}
	}()
	select {
	case <-time.After(timeout):
		logger.Warn("DownloadService.DownloadFeed: Timeout reached when downloading ", url)
		return downloader.Response{}, errors.New("Timeout reached")
	case p := <-ch:
		return p.Response, p.error
	}
	panic("unreachable")
}
