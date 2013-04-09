package coverage

import (
	"git.300brand.com/coverage/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Feed struct {
	ID           bson.ObjectId `bson:"_id"`
	ObjectId     uint64        // Comes from the bridge package when inserting from the frontend
	QueueId      uint64        // Comes from the bridge package when inserting from the frontend
	URL          *url.URL
	Disabled     bool
	Content      []byte     `bson:"-"`
	Articles     []*Article `bson:"-"` // Temporary Article storage
	URLs         []*url.URL
	Log          logger.Entries
	Added        time.Time
	Updated      time.Time
	LastDownload time.Time
}

func NewFeed() (f *Feed) {
	f = &Feed{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	return
}

func (f *Feed) AddURL(u *url.URL) bool {
	urlCopy := *u

	s := urlCopy.String()
	for _, v := range f.URLs {
		if v.String() == s {
			return false
		}
	}
	f.URLs = append(f.URLs, &urlCopy)
	return true
}

func (f *Feed) Files() []File {
	return []File{
		{
			Name:        f.ID.Hex() + ".xml",
			ContentType: "text/xml",
			Data:        f.Content,
		},
	}
}

func (f *Feed) Downloaded() {
	f.LastDownload = time.Now()
}
