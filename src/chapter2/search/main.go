package main

import (
	"log"
	"os"
	"sync"

	"github.com/goinaction/code/src/chapter2/search/data"
	search "github.com/goinaction/code/src/chapter2/search/feed"
	"github.com/goinaction/code/src/chapter2/search/rss"
)

// init is called before main.
func init() {
	// Change the log to write to stdout instead
	// of its default device stderr.
	log.SetOutput(os.Stdout)
}

// NewMatcher is a factory that creates matcher values based
// on the type of feed specified.
func NewMatcher(feed data.Feed) search.Matcher {
	// Create the right type of matcher for this search.
	switch feed.Type {
	case "rss":
		return rss.NewMatcher(&feed)

		// TODO: Add new Matchers here
	}

	log.Fatalln("Invalid Feed Type")
	return nil
}

// main is the entry point for the program.
func main() {
	// Search term we are looking for.
	searchTerm := "president"

	// Load the feeds from the data file.
	feeds, err := data.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Create a channel to receive the results on.
	results := make(chan search.Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		// Create a matcher for the search.
		matcher := NewMatcher(feed)

		// Launch the goroutine to perform the search.
		go func() {
			search.Search(matcher, searchTerm, results)
			waitGroup.Done()
		}()
	}

	// Launch a goroutine so we can shutdown the program
	// once the last feed sends its results.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel.
		close(results)
	}()

	// Display the results.
	search.Display(results)
}
