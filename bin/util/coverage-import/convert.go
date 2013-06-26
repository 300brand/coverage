package main

import (
	"compress/gzip"
	"fmt"
	"git.300brand.com/coverage"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/url"
	"os"
	"time"
)

var feedIdCache = map[uint64]bson.ObjectId{}

func ConvertArticle(in Article) (out *coverage.Article, err error) {
	out = coverage.NewArticle()

	start := time.Now()
	out.Text.HTML, err = ArticleHTML(in.Id)
	log.Printf("         Read HTML %s", time.Now().Sub(start))

	if out.URL, err = url.Parse(string(in.Url)); err != nil {
		return
	}
	if out.Added, err = time.ParseInLocation(timeLayout, in.Added, loc); err != nil {
		return
	}
	if out.Published, err = time.ParseInLocation(timeLayout, in.Published, loc); err != nil {
		out.Published = time.Time{}
	}
	if out.FeedId, err = FeedId(in.FeedId); err != nil {
		return
	}

	return
}

func FeedId(id uint64) (fid bson.ObjectId, err error) {
	fid, ok := feedIdCache[id]
	if ok {
		return
	}

	f, err := conn.Mongo.GetFeed(bson.M{"feedid": id})
	if err != nil {
		if err.Error() == "not found" {
			err = fmt.Errorf("Could not find feed: %d", id)
		}
		return
	}
	fid = f.ID
	feedIdCache[id] = fid
	return
}

func ArticleHTML(id uint64) ([]byte, error) {
	f, err := os.Open(IdPath(id))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}

func IdPath(id uint64) string {
	dir := ((id + 999) / 1000) * 1000
	return fmt.Sprintf("%s/%d/%d.gz", path, dir, id)
}
