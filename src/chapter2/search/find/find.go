package find

import "log"

type (
	// Result contains the result of a search.
	Result struct {
		Field   string
		Content string
	}

	// Matcher defines the behavior required by the Search function.
	Matcher interface {
		Match(searchTerm string) ([]Result, error)
	}
)

// Search pulls down each feed looking for the search term.
func Search(matcher Matcher, searchTerm string, result chan Result) {
	// Search the data for the search term.
	searchResults, err := matcher.Match(searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, searchResult := range searchResults {
		result <- searchResult
	}
}
