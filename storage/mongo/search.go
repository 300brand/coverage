package mongo

import (
	"git.300brand.com/coverage"
	"labix.org/v2/mgo/bson"
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
