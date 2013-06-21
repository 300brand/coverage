package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/downloader"
	"github.com/skynetservices/skynet"
)

func (s *ArticleDownload) Download(ri *skynet.RequestInfo, in *coverage.Article, out *coverage.Article) (err error) {
	if err = downloader.Article(in); err != nil {
		return
	}
	*out = *in
	return
}
