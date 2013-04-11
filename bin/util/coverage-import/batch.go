package main

import (
	"log"
)

func GetBatch(lastId, limit uint64) (batch []Article, err error) {
	batch = make([]Article, 0, limit)
	rows, err := conn.ArticleStmt.Query(lastId, limit)
	if err != nil {
		return
	}
	for rows.Next() {
		a := Article{}
		err = rows.Scan(&a.Id, &a.FeedId, &a.Title, &a.Url, &a.Published, &a.Added)
		if err != nil {
			return
		}
		batch = append(batch, a)
	}
	err = rows.Err()
	return
}

func ProcessBatch(batch []Article, ch chan interface{}) (newStart uint64) {
	for _, b := range batch {
		newStart = b.Id
		go func(in Article) {
			a, err := ConvertArticle(in)
			if err != nil {
				log.Print(err)
				ch <- err
				return
			}
			for _, s := range services {
				if err := s.Update(a); err != nil {
					ch <- err
					return
				}
			}
			ch <- a
		}(b)
	}
	return
}
