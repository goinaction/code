package find

import (
	"log"
	"sync"

	"github.com/goinaction/code/src/chapter2/search/feeds"
)

type (
	// Result contains the result of a search.
	Result struct {
		Field   string
		Content string
	}

	// Matcher defines the behavior required by the Search function.
	Matcher interface {
		Match(site *feeds.Site, searchTerm string) ([]Result, error)
	}
)

// Search pulls down each feed looking for the search term.
func Search(matcher Matcher, searchTerm string, site feeds.Site, result chan Result, waitGroup *sync.WaitGroup) {
	// Call done so we can report we are finished processing.
	defer func() {
		waitGroup.Done()
	}()

	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", site.Type, site.Name, site.Uri)

	// Search the data for the search term.
	searchResults, err := matcher.Match(&site, searchTerm)
	if err != nil {
		log.Printf("%s : %s", site.Uri, err)
		return
	}

	// Write the results to the channel.
	for _, searchResult := range searchResults {
		result <- searchResult
	}
}
