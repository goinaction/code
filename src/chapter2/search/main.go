package main

import (
	"log"

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
	results := make(chan []find.Result)

	// Launch a goroutine for each feed to find the results.
	for _, site := range sites {
		var matcher find.Matcher

		// Create the right type of matcher for this search.
		switch site.Type {
		case "rss":
			matcher = new(rss.Search)

		default:
			log.Fatalln("Invalid Type")
		}

		// Launch the goroutine to perform the search.
		go find.Search(matcher, searchTerm, site, results)
	}

	// Wait for the result from each goroutine
	for site := 0; site < len(sites); site++ {
		select {
		case found := <-results:
			for _, result := range found {
				log.Printf("%s:\n%s\n\n", result.Field, result.Content)
			}
		}
	}
}
