package main

import (
	"log"
	"sync"

	"github.com/goinaction/code/src/chapter2/search/feeds"
	"github.com/goinaction/code/src/chapter2/search/find"
	"github.com/goinaction/code/src/chapter2/search/rss"
)

// main is the entry point for the program.
func main() {
	// Search term we are looking for.
	searchTerm := "president"

	// Load the feeds for the data file.
	sites, err := feeds.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Perpare the slice of results and the channel to
	// retrieve the results on.
	result := make(chan find.Result)

	// Wait for the result from each goroutine
	go func() {
		for found := range result {
			log.Printf("%s:\n%s\n\n", found.Field, found.Content)
		}
	}()

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(sites))

	// Launch a goroutine for each feed to find the results.
	for _, site := range sites {
		var matcher find.Matcher

		// Create the right type of matcher for this search.
		switch site.Type {
		case "rss":
			matcher = &rss.Search{
				Site: site,
			}

		default:
			log.Fatalln("Invalid Type")
		}

		// Launch the goroutine to perform the search.
		go find.Search(matcher, searchTerm, result, &waitGroup)
	}

	// Wait for everything to be processed.
	waitGroup.Wait()

	// Close the channel and exit.
	close(result)
}
