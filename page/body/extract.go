package body

import (
	"github.com/moovweb/gokogiri"
)

func GetBody(b []byte) (body []byte, err error) {
	doc, err := gokogiri.ParseHtml(b)
	if err != nil {
		return
	}
	doc.Root()
	return
}
