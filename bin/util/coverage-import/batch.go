package main

import (
	"log"
)

func GetBatch(lastId uint64, batch []Article) (err error) {
	log.Printf("lastId: %d limit: %d", lastId, cap(batch))
	rows, err := conn.ArticleStmt.Query(lastId, cap(batch))
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
	log.Printf("Batch size: %d", len(batch))
	return rows.Err()
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
