package body

import (
	"errors"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
	"github.com/jbaikge/logger"
)

type ArticleService struct{}

var _ service.ArticleService = ArticleService{}

func NewArticleService() ArticleService {
	return ArticleService{}
}

func (s ArticleService) Update(a *coverage.Article) (err error) {
	a.Log.Service("body.ArticleService")
	return SetBody(a)
}

func SetBody(a *coverage.Article) (err error) {
	logger.Trace.Printf("SetBody: called %s", a.ID.Hex())
	if a.Text.HTML == nil {
		err = errors.New("HTML not set, did you run the downloader service?")
		logger.Error.Printf("SetBody: A %s %s", a.ID.Hex(), err)
		return
	}
	cleaned, err := CleanHTML(a.Text.HTML)
	if err != nil {
		logger.Error.Printf("SetBody: A %s %s", a.ID.Hex(), err)
		return
	}
	if a.Text.Body, err = GetBody(cleaned); err != nil {
		logger.Error.Printf("SetBody: A %s %s", a.ID.Hex(), err)
		return
	}
	logger.Debug.Printf("SetBody: A %s Updated body", a.ID.Hex())
	return
}
