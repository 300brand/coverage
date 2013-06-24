package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/article/body"
	"github.com/skynetservices/skynet"
)

func (s *ArticleBody) Process(ri *skynet.RequestInfo, in *coverage.Article, out *coverage.Article) (err error) {
	if err = body.SetBody(in); err != nil {
		return
	}
	*out = *in
	return
}
