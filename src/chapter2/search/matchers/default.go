package matchers

import (
	"github.com/goinaction/code/src/chapter2/search/match"
)

// defaultMatcher implements the Matcher interface.
type defaultMatcher struct {
	*match.Feed
}

// Search looks at the document for the specified search term.
func (m *defaultMatcher) Search(searchTerm string) ([]*match.Result, error) {
	return nil, nil
}
