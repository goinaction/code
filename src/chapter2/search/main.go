package main

import (
	"log"

	"github.com/goinaction/code/src/chapter2/search/feeds"
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
	captureResults := make(chan []rss.SearchResult)

	// Launch goroutines to find us our results.
	for _, site := range sites {
		go find(searchTerm, site, captureResults)
	}

	// Wait for the result from each goroutine
	for site := 0; site < len(sites); site++ {
		select {
		case findResults := <-captureResults:
			for _, result := range findResults {
				display(&result)
			}
		}
	}
}

// find pulls down each feed and searches for the results.
func find(searchTerm string, site feeds.Site, captureResults chan []rss.SearchResult) {
	// Make sure each find returns a result.
	var err error
	defer func() {
		if err != nil {
			captureResults <- nil
		}
	}()

	log.Printf("Search Feed Site[%s] For Uri[%s]\n", site.Name, site.Uri)

	// Retrieve the RSS feed document.
	document, err := rss.Retrieve(site.Uri)
	if err != nil {
		log.Printf("%s : %s", site.Uri, err)
		return
	}

	// Search the document for the search term.
	searchResults, err := rss.Search(document, searchTerm)
	if err != nil {
		log.Printf("%s : %s", site.Uri, err)
		return
	}

	// Write the results to the channel.
	captureResults <- searchResults
}

// display logs the results of the serach to the console.
func display(result *rss.SearchResult) {
	switch result.Field {
	case "Title":
		log.Printf("Title:\n%s\n\n", result.Document.Title)
	case "Description":
		log.Printf("Description:\n%s\n\n", result.Document.Description)
	}
}
