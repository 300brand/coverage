package main

import (
	"log"
)

func GetBatch(lastId uint64, batch []Article) (n int, err error) {
	log.Printf("lastId: %d limit: %d", lastId, len(batch))
	rows, err := conn.ArticleStmt.Query(lastId, len(batch))
	if err != nil {
		return
	}
	for n = 0; rows.Next(); n++ {
		a := Article{}
		err = rows.Scan(&a.Id, &a.FeedId, &a.Title, &a.Url, &a.Published, &a.Added)
		if err != nil {
			return
		}
		batch[n] = a
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
