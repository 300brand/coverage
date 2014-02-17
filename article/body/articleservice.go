package body

import (
	"errors"
	"fmt"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
	"github.com/300brand/logger"
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

	var cleaned []byte
	var body coverage.Body
	pages := make([][]byte, 1, len(a.Text.Pages)+1)
	pages[0] = a.Text.HTML
	pages = append(pages, a.Text.Pages...)

	for i, html := range pages {
		cleaned, err = CleanHTML(html)
		if err != nil {
			logger.Error.Printf("SetBody: A %s %s", a.ID.Hex(), err)
			return
		}
		if body, err = GetBody(cleaned); err != nil {
			logger.Error.Printf("SetBody: A %s %s", a.ID.Hex(), err)
			return
		}

		a.Text.Body.HTML = append(a.Text.Body.HTML, []byte(fmt.Sprintf(`<div class="page-%d">`, i+1))...)
		a.Text.Body.HTML = append(a.Text.Body.HTML, body.HTML...)
		a.Text.Body.HTML = append(a.Text.Body.HTML, []byte(`</div>`)...)

		a.Text.Body.Text = append(a.Text.Body.Text, body.Text...)
		a.Text.Body.Text = append(a.Text.Body.Text, '\n')
	}
	logger.Debug.Printf("SetBody: A %s Updated body", a.ID.Hex())
	return
}
