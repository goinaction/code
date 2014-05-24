package match

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
	Search(searchTerm string) ([]*Result, error)
}

// Match pulls down each feed looking for the search term.
func Match(matcher Matcher, searchTerm string, results chan<- *Result) {
	// Search the data for the search term.
	searchResults, err := matcher.Search(searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the indivdual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is close the for loop terminates.
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
