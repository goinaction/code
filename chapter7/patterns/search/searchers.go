// Copyright Information.

// Package search : seachers.go contains all the different implementations
// for the existing searchers.
package search

import (
	"log"
	"math/rand"
	"time"
)

// init is called before main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Google provides support for Google searches.
type Google struct{}

// NewGoogle returns a Google Searcher value.
func NewGoogle() Searcher {
	return Google{}
}

// Search implements the Searcher interface. It performs a search
// against Google.
func (g Google) Search(searchTerm string, searchResults chan<- []Result) {
	log.Printf("Google : Search : Started : searchTerm[%s]\n", searchTerm)

	// Slice for the results.
	var results []Result

	// On return send the results we have.
	defer func() {
		log.Println("Google : Search : Info : Sending Results")
		searchResults <- results
	}()

	// Simulate an amount of time for the search.
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// Simulate a result for the search.
	results = append(results, Result{
		Engine:      "Google",
		Title:       "The Go Programming Language",
		Description: "The Go Programming Language",
		Link:        "https://golang.org/",
	})

	log.Printf("Google : Search : Completed : Found[%d]\n", len(results))
}

// Bing provides support for Bing searches.
type Bing struct{}

// NewBing returns a Bing Searcher value.
func NewBing() Searcher {
	return Bing{}
}

// Search implements the Searcher interface. It performs a search
// against Bing.
func (b Bing) Search(searchTerm string, searchResults chan<- []Result) {
	log.Printf("Bing : Search : Started : searchTerm[%s]\n", searchTerm)

	// Slice for the results.
	var results []Result

	// On return send the results we have.
	defer func() {
		log.Println("Bing : Search : Info : Sending Results")
		searchResults <- results
	}()

	// Simulate an amount of time for the search.
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// Simulate a result for the search.
	results = append(results, Result{
		Engine:      "Bing",
		Title:       "A Tour of Go",
		Description: "Welcome to a tour of the Go programming language.",
		Link:        "http://tour.golang.org/",
	})

	log.Printf("Bing : Search : Completed : Found[%d]\n", len(results))
}

// Yahoo provides support for Yahoo searches.
type Yahoo struct{}

// NewYahoo returns a Yahoo Searcher value.
func NewYahoo() Searcher {
	return Yahoo{}
}

// Search implements the Searcher interface. It performs a search
// against Yahoo.
func (y Yahoo) Search(searchTerm string, searchResults chan<- []Result) {
	log.Printf("Yahoo : Search : Started : searchTerm[%s]\n", searchTerm)

	// Slice for the results.
	var results []Result

	// On return send the results we have.
	defer func() {
		log.Println("Yahoo : Search : Info : Sending Results")
		searchResults <- results
	}()

	// Simulate an amount of time for the search.
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// Simulate a result for the search.
	results = append(results, Result{
		Engine:      "Yahoo",
		Title:       "Go Playground",
		Description: "The Go Playground is a web service that runs on golang.org's servers",
		Link:        "http://play.golang.org/",
	})

	log.Printf("Yahoo : Search : Completed : Found[%d]\n", len(results))
}
