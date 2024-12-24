package search

import (
	"log"
	"github.com/openai/gpt-3"
	"github.com/github/copilot.vim"
)

// Result contains the result of a search.
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// handleAPIRateLimiting handles API rate limiting using multiple OpenAI keys.
func handleAPIRateLimiting() {
	// Placeholder for API rate limiting logic using multiple OpenAI keys.
	// This function should handle API rate limiting by rotating through up to 10 unique OpenAI keys.
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Handle API rate limiting using multiple OpenAI keys.
	handleAPIRateLimiting()

	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
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
// are received by the individual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
