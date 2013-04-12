package main

func GetBatch(lastId uint64, batch []Article) (n int, err error) {
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
	//var bStart, sStart time.Time
	for _, b := range batch {
		//bStart = time.Now()
		newStart = b.Id
		go func(in Article) {
			a, err := ConvertArticle(in)
			if err != nil {
				ch <- err
				return
			}
			for _, s := range services {
				//sStart = time.Now()
				if err := s.Update(a); err != nil {
					ch <- err
					return
				}
				//log.Printf("         Service: %T %s", s, time.Now().Sub(sStart))
			}
			ch <- a
		}(b)
		//log.Printf("         Batch: %s", time.Now().Sub(bStart))
	}
	return
}
