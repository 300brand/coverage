package mongo

import (
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/article/lexer"
	"git.300brand.com/coverage/search"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"sync"
	"time"
)

const (
	SearchCollection        = "Search"
)

// Add a batch of search results from a date
func (m *Mongo) AddSearchResults(id bson.ObjectId, articles []bson.ObjectId) (err error) {
	change := bson.M{
		"$inc": bson.M{
			"daysleft": -1,
			"results":  len(articles),
		},
		"$addToSet": bson.M{
			"articles": bson.M{
				"$each": articles,
			},
		},
	}
	if err = m.C.Search.UpdateId(id, change); err != nil {
		return
	}

	complete := bson.M{
		"$set": bson.M{
			"complete": time.Now(),
		},
	}
	return m.C.Search.Update(bson.M{"_id": id, "daysleft": 0}, complete)
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
	log := log.New(os.Stdout, fmt.Sprintf("%d ", dtoi(t)), log.Lmicroseconds|log.Lshortfile)
	var (
		wg         sync.WaitGroup
		terms      = lexer.Keywords([]byte(query))
		boolSearch = search.NewBoolean(query)
		idFilter   = search.NewIdFilter(boolSearch.MinTerms())
		idChan     = make(chan bson.ObjectId, 1)
		saveChan   = make(chan bson.ObjectId, 1)
		results    = make([]bson.ObjectId, 0, 10)
	)
	idFilter.UseChan()

	// Collect Article IDs from Keyword objects
	go func() {
		for id := range idChan {
			log.Printf("Got %s from idChan", id)
			idFilter.Add(id)
		}
		log.Println("Closing idFilter.Chan")
		close(idFilter.Chan)
	}()

	// Get each article from the idFilter to run fulltext analysis
	go func() {
		var wg sync.WaitGroup
		for id := range idFilter.Chan {
			log.Printf("Got %s from idFilter.Chan", id)
			wg.Add(1)
			go func(id bson.ObjectId) {
				defer wg.Done()
				defer log.Println("Finished article processing")

				// Fetch Article
				a, err := m.GetArticle(bson.M{"_id": id})
				if err != nil {
					// TODO Flag err != Not found
					return
				}

				// Boolean
				if !boolSearch.Match(a.Text.Body.Text) {
					return
				}

				// Send to save result
				log.Printf("Sending %s along saveChan", id)
				saveChan <- id
			}(id)
		}
		wg.Wait()
		close(saveChan)
	}()

	// Query for each keyword for this date
	for _, term := range terms {
		wg.Add(1)
		go func(term string) {
			defer wg.Done()
			defer log.Printf("Finished keyword query for %s", term)

			id := &coverage.KeywordId{
				Date:    t,
				Keyword: term,
			}
			kw := &coverage.Keyword{}
			log.Printf("Querying: %v", id)
			if err := m.GetKeyword(id, kw); err != nil {
				log.Printf("Query error in Mongo.Keyword: %s", err)
				return
			}
			for _, id := range kw.Articles {
				log.Printf("Sending %s", id)
				idChan <- id
			}
		}(term)
	}
	log.Println("Waiting for keyword queries to finish")
	wg.Wait()
	close(idChan)
	log.Println("Closed idChan")

	for id := range saveChan {
		results = append(results, id)
	}
	log.Printf("Finished GetKeywords")

	if err = m.AddSearchResults(searchId, results); err != nil {
		log.Printf("Error Adding Search Results: %s", err)
		return
	}

	return
}
