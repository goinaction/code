package main

import (
	"log"
	"sync"
)

// NewMatcher is a factory that creates matcher values based
// on the type of feed specified.
func NewMatcher(feed Feed) Matcher {
	// Create the right type of matcher for this search.
	switch feed.Type {
	case "rss":
		return &rssMatcher{Feed: feed}

		// TODO: Add new Matchers here
	}

	// TODO: make a matcher implementation that never matches anything
	// then return that.
	// Right now your program will do this when you hit an atom feed
	// "Invalid Feed Type"
	// Panic: call to nil interface method.
	log.Fatal("Invalid Feed Type")
	return nil
}

// main is the entry point for the program.
func main() {
	// Search term we are looking for.
	searchTerm := "president"

	// Load the feeds from the data file.
	feeds, err := Load()
	if err != nil {
		log.Fatal(err)
	}

	// Create a channel to receive the results on.
	results := make(chan Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// TODO: needs a comment why this is done here rather than inside the loop.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		// Create a matcher for the search.
		matcher := NewMatcher(feed)

		// Launch the goroutine to perform the search.
		go func() {
			Search(matcher, searchTerm, results)
			waitGroup.Done()
		}()
	}

	// TODO: this is correct, but needs an explanation, ie
	// why not write
	// waitGroup.Wait() ; close(results); search.Display(results)
	// or
	// go search.Display(results); waitGroup.Wait(); close(results)
	//
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel.
		close(results)
	}()

	// Display the results.
	Display(results)
}
