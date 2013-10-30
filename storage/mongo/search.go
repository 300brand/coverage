package mongo

import (
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/article/lexer"
	"github.com/300brand/coverage/search"
	"github.com/300brand/logger"
	"labix.org/v2/mgo/bson"
	"sync"
	"time"
)

const SearchCollection = "Search"

// Add a batch of search results from a date
func (m *Mongo) AddSearchResults(id bson.ObjectId, articles []bson.ObjectId) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("AddSearchResults called")
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
	if err = c.Search.UpdateId(id, change); err != nil {
		logger.Error.Printf("AddSearchResults: UpdateId - %s", err)
		return
	}

	complete := bson.M{
		"$set": bson.M{
			"complete": time.Now(),
		},
	}
	c.Search.Update(bson.M{"_id": id, "daysleft": 0}, complete)
	return
}

func (m *Mongo) GetSearch(id bson.ObjectId, s *coverage.Search) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("GetSearch called")
	return c.Search.FindId(id).One(s)
}

func (m *Mongo) UpdateSearch(s *coverage.Search) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("UpdateSearch called")
	_, err = c.Search.UpsertId(s.Id, s)
	return
}

func (m *Mongo) CompileResults(id bson.ObjectId) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("CompileResults called")
	s := &coverage.Search{}
	if err = m.GetSearch(id, s); err != nil {
		return
	}
	if err = c.SearchResults.Find(bson.M{"searchid": id}).All(&s.Articles); err != nil {
		return
	}
	return m.UpdateSearch(s)
}

func (m *Mongo) DateSearch(searchId bson.ObjectId, query string, t time.Time) (err error) {
	c := m.Copy()
	defer c.Close()

	logger.Trace.Printf("DateSearch called")
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
			logger.Debug.Printf("Got %s from idChan", id)
			idFilter.Add(id)
		}
		logger.Debug.Println("Closing idFilter.Chan")
		close(idFilter.Chan)
	}()

	// Get each article from the idFilter to run fulltext analysis
	go func() {
		var wg sync.WaitGroup
		for id := range idFilter.Chan {
			logger.Debug.Printf("Got %s from idFilter.Chan", id)
			wg.Add(1)
			go func(id bson.ObjectId) {
				defer wg.Done()
				defer logger.Debug.Println("Finished article processing")

				// Fetch Article
				a := &coverage.Article{}
				if err := m.GetArticle(bson.M{"_id": id}, a); err != nil {
					// TODO Flag err != Not found
					return
				}

				// Boolean
				if !boolSearch.Match(a.Text.Body.Text) {
					return
				}

				// Send to save result
				logger.Debug.Printf("Sending %s along saveChan", id)
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
			defer logger.Debug.Printf("Finished keyword query for %s", term)

			id := &coverage.KeywordId{
				Date:    t,
				Keyword: term,
			}
			kw := &coverage.Keyword{}
			logger.Debug.Printf("Querying: %v", id)
			if err := m.GetKeyword(id, kw); err != nil {
				logger.Debug.Printf("Query error in Mongo.Keyword: %s", err)
				return
			}
			for _, id := range kw.Articles {
				logger.Debug.Printf("Sending %s", id)
				idChan <- id
			}
		}(term)
	}
	logger.Debug.Println("Waiting for keyword queries to finish")
	wg.Wait()
	close(idChan)
	logger.Debug.Println("Closed idChan")

	for id := range saveChan {
		results = append(results, id)
	}
	logger.Debug.Printf("Finished GetKeywords")

	if err = m.AddSearchResults(searchId, results); err != nil {
		logger.Debug.Printf("Error Adding Search Results: %s", err)
		return
	}

	return
}
