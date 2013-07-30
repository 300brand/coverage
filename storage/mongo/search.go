package mongo

import (
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/article/lexer"
	"git.300brand.com/coverage/search"
	"labix.org/v2/mgo/bson"
	"log"
	"sync"
	"time"
)

const (
	SearchCollection        = "Search"
	SearchResultsCollection = "SearchResults"
)

// Add a batch of search results from a date
func (m *Mongo) AddSearchResults(id bson.ObjectId, results []*coverage.SearchResult) (err error) {
	change := bson.M{
		"$inc": bson.M{
			"daysleft": -1,
			"results":  len(results),
		},
	}
	if err = m.C.Search.UpdateId(id, change); err != nil {
		return
	}
	for _, r := range results {
		if err = m.C.SearchResults.Insert(r); err != nil {
			return
		}
	}
	return
}

func (m *Mongo) GetSearch(id bson.ObjectId, s *coverage.Search) (err error) {
	return m.C.Search.FindId(id).One(s)
}

func (m *Mongo) UpdateSearch(s *coverage.Search) (err error) {
	_, err = m.C.Search.UpsertId(s.Id, s)
	return
}

func (m *Mongo) CompileResults(id bson.ObjectId) (err error) {
	s := &coverage.Search{}
	if err = m.GetSearch(id, s); err != nil {
		return
	}
	if err = m.C.SearchResults.Find(bson.M{"searchid": id}).All(&s.Articles); err != nil {
		return
	}
	return m.UpdateSearch(s)
}

func (m *Mongo) DateSearch(searchId bson.ObjectId, query string, t time.Time) (err error) {
	var wg sync.WaitGroup

	terms := lexer.Keywords([]byte(query))
	boolSearch := search.NewBoolean(query)
	idChan := make(chan bson.ObjectId)
	idFilter := search.NewIdFilter(boolSearch.MinTerms())
	idFilter.UseChan()

	// Query for each keyword for this date
	for _, term := range terms {
		wg.Add(1)
		go func(term string) {
			defer wg.Done()

			id := &coverage.KeywordId{
				Date:    t,
				Keyword: term,
			}
			kw := &coverage.Keyword{}
			log.Printf("Querying: %v", id)
			if err := m.GetKeyword(id, kw); err != nil {
				log.Printf("Query error in Mongo.Keyword(%v): %s", id, err)
				return
			}
			for _, id := range kw.Articles {
				idFilter.Add(id)
			}
		}(term)
	}

	// Wait for the keyword queries to finish
	go func() {
		wg.Wait()
		close(idFilter.Chan)
	}()

	// Pull Articles for body matches
	for id := range idFilter.Chan {
		wg.Add(1)
		go func(id bson.ObjectId) {
			defer wg.Done()

			a, err := m.GetArticle(bson.M{"_id": id})
			if err != nil {
				log.Printf("Mongo.DateSearch.GetArticle: %s", err)
				return
			}
			if boolSearch.Match(a.Text.Body.Text) {
				idChan <- id
			}
		}(id)
	}

	go func() {
		wg.Wait()
		close(idChan)
	}()

	ids := make([]*coverage.SearchResult, 0, 10)
	for id := range idChan {
		ids = append(ids, &coverage.SearchResult{
			SearchId:  searchId,
			ArticleId: id,
		})
	}

	log.Println("Finished GetKeywords")

	return m.AddSearchResults(searchId, ids)
}
