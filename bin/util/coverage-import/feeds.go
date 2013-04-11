package main

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/service"
	"github.com/skynetservices/mgo/bson"
	"log"
	"net/url"
	"time"
)

type Feed struct {
	Id       uint64
	ObjectId uint64
	Url      string
	Added    string
}

var fServices []service.FeedService

func ImportFeeds() (err error) {
	q := `
		SELECT id, IFNULL(remote_id, 0), url, _added
		FROM Feed
		WHERE
			url NOT LIKE 'https://www.googleapis.com/%'
			AND (
				expires IS NULL
				OR expires > NOW()
			)
			AND failures < 14
		ORDER BY id
	`
	rows, err := conn.MySQL.Query(q)
	if err != nil {
		return
	}

	in := Feed{}
	for rows.Next() {
		if err := rows.Scan(&in.Id, &in.ObjectId, &in.Url, &in.Added); err != nil {
			return err
		}
		log.Printf("[%04d] %s", in.Id, in.Url)
		f, err := ConvertFeed(in)
		if err != nil {
			return err
		}
		for _, s := range fServices {
			if err := s.Update(f); err != nil {
				return err
			}
		}
	}
	return
}

func ConvertFeed(in Feed) (out *coverage.Feed, err error) {
	out = coverage.NewFeed()

	out.FeedId = in.Id
	out.ObjectId = in.ObjectId

	if out.URL, err = url.Parse(in.Url); err != nil {
		return
	}
	if out.Added, err = time.ParseInLocation(timeLayout, in.Added, loc); err != nil {
		return
	}

	out.ID = bson.NewObjectIdWithTime(out.Added)

	return
}
