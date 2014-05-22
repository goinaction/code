package main

import (
	"log"
	"sync"

	"github.com/goinaction/code/src/chapter2/search/data"
	"github.com/goinaction/code/src/chapter2/search/feed"
	"github.com/goinaction/code/src/chapter2/search/rss"
)

// main is the entry point for the program.
func main() {
	// Search term we are looking for.
	searchTerm := "president"

	// Load the feeds from the data file.
	feeds, err := data.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Start the display routine.
	results := feed.Display()

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, data := range feeds {
		var matcher feed.Matcher

		// Create the right type of matcher for this search.
		switch data.Type {
		case "rss":
			matcher = rss.NewMatcher(data)

		default:
			log.Fatalln("Invalid Type")
		}

		// Launch the goroutine to perform the search.
		go func() {
			defer waitGroup.Done()
			feed.Search(matcher, searchTerm, results)
		}()
	}

	// Wait for everything to be processed.
	waitGroup.Wait()

	// Close the channel and exit.
	close(results)
}
