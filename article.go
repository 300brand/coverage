package coverage

import (
	"git.300brand.com/coverage/article/body"
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Article struct {
	ID          bson.ObjectId
	Title       string
	URL         url.URL
	ProperNames ProperNames
	Body        body.Body
	Logs        logger.Entries
	Times       struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
		Published time.Time
	}
}
