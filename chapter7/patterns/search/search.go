// Copyright Information.

// Package search : search.go manages the searching of results against Google, Yahoo and Bing.
package search

import "log"

// Result represents a search result that was found.
type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

// Searcher declares an interface used to leverage different
// search engines to find results.
type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	s.searchers["google"] = NewGoogle()
}

func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = NewBing()
}

func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = NewYahoo()
}

func First(s *searchSession) { s.first = true }

// Submit uses goroutines and channels to perform a search against the three
// leading search engines concurrently.
func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)

	for _, opt := range options {
		opt(&session)
	}

	var result []Result

	// Perform the searches concurrently. Using a map because
	// it returns the searchers in a random order every time.
	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	// Wait for the results to come back.
	for search := 0; search < len(session.searchers); search++ {
		// If we just want the first result, don't wait any longer by
		// concurrently discarding the remaining searchResults.
		// Failing to do so will leave the Searchers blocked forever.
		if session.first && search > 0 {
			go func() {
				result = append(result, <-session.resultChan...)
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(result))
			}()
			continue
		}

		// Wait to recieve results.
		log.Println("search : Submit : Info : Waiting For Results...")
		result = append(result, <-session.resultChan...)

		// Save the results to the final slice.
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(result))
	return result
}
