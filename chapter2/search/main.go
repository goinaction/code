package main

import (
	"log"
	"sync"

	"github.com/goinaction/code/src/chapter2/search/match"
	"github.com/goinaction/code/src/chapter2/search/matchers"
)

// main is the entry point for the program.
func main() {
	// Search term we are looking for.
	searchTerm := "president"

	// RetrieveFeed returns the list of feeds to search through.
	feeds, err := match.RetrieveFeed()
	if err != nil {
		log.Fatal(err)
	}

	// Create a channel to receive match results to display.
	results := make(chan *match.Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		// Create a matcher for the search.
		matcher := matchers.NewMatcher(feed)

		// Launch the goroutine to perform the search.
		go func() {
			match.Match(matcher, searchTerm, results)
			waitGroup.Done()
		}()
	}

	// Launch a groutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are avaiable and
	// return after the final result is displayed.
	match.Display(results)
}
