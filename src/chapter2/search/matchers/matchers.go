package matchers

import (
	"github.com/goinaction/code/src/chapter2/search/match"
)

// NewMatcher is a factory that creates matcher values based
// on the type of feed specified.
func NewMatcher(feed *match.Feed) match.Matcher {
	// Create the right type of matcher for this search.
	switch feed.Type {
	case "rss":
		return &rssMatcher{feed}

	default:
		return &defaultMatcher{feed}
	}
}
