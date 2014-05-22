package find

import (
	"log"
	"sync"

	"github.com/goinaction/code/src/chapter2/search/feeds"
)

type (
	// SearchData contains the data we can search on.
	SearchData struct {
		Title       string
		Description string
	}

	// Result contains the result of a search.
	Result struct {
		Field   string
		Content string
	}
)

type (
	// Matcher defines the behavior required by the Search function.
	Matcher interface {
		Retrieve(site feeds.Site) ([]SearchData, error)
		Match(document []SearchData, searchTerm string) ([]Result, error)
	}
)

// Search pulls down each feed looking for the search term.
func Search(matcher Matcher, searchTerm string, site feeds.Site, result chan Result, waitGroup *sync.WaitGroup) {
	// Call done so we can report we are finished processing.
	defer func() {
		waitGroup.Done()
	}()

	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", site.Type, site.Name, site.Uri)

	// Retrieve the search data.
	searchData, err := matcher.Retrieve(site)
	if err != nil {
		log.Printf("%s : %s", site.Uri, err)
		return
	}

	// Search the data for the search term.
	searchResults, err := matcher.Match(searchData, searchTerm)
	if err != nil {
		log.Printf("%s : %s", site.Uri, err)
		return
	}

	// Write the results to the channel.
	for _, searchResult := range searchResults {
		result <- searchResult
	}
}
