package coverage

import (
	"github.com/300brand/coverage/logger"
	"github.com/300brand/coverage/validurl"
	log "github.com/300brand/logger"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Feed struct {
	ID            bson.ObjectId `bson:"_id"`
	PublicationId bson.ObjectId `json:",omitempty"`
	URL           *url.URL
	Deleted       bool
	Content       []byte
	Articles      []*Article // Temporary Article storage; cleared before each save
	URLs          []*url.URL
	Log           logger.Entries
	Added         time.Time
	Updated       time.Time
	LastDownload  time.Time `bson:",omitempty"`
}

func NewFeed() (f *Feed) {
	f = &Feed{
		ID:    bson.NewObjectId(),
		Added: time.Now(),
	}
	return
}

func (f *Feed) AddURL(u *url.URL) bool {
	if !validurl.IsValid(u) {
		log.Debug.Printf("[P:%s] [F:%s] [U:%s] Invalid URL. Skipping.", f.PublicationId, f.ID, u)
		return false
	}

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

// TODO remove file storage
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
