package main

import (
	"fmt"
	"log"
)

// Result contains the result of a search.
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by the Search function.
type Matcher interface {
	Match(searchTerm string) ([]Result, error)
}

// Search pulls down each feed looking for the search term.
func Search(matcher Matcher, searchTerm string, results chan<- Result) {
	// Search the data for the search term.
	searchResults, err := matcher.Match(searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, searchResult := range searchResults {
		results <- searchResult
	}
}

// Display writes results to the console window.
func Display(results chan Result) {
	// Wait for results from the different feeds and
	// display them.
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
