package find

import (
	"log"
	"regexp"
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
	searchResults, err := match(searchData, searchTerm)
	if err != nil {
		log.Printf("%s : %s", site.Uri, err)
		return
	}

	// Write the results to the channel.
	for _, searchResult := range searchResults {
		result <- searchResult
	}
}

// match looks at the document for the specified search term.
func match(searchData []SearchData, searchTerm string) ([]Result, error) {
	var results []Result

	for _, data := range searchData {
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, data.Title)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, Result{
				Field:   "Title",
				Content: data.Title,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, data.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, Result{
				Field:   "Description",
				Content: data.Description,
			})
		}
	}

	return results, nil
}
