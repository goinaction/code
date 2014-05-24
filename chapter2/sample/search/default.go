package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registeres the default matcher with the program.
func init() {
	Register("default", &defaultMatcher{})
}

// Search implements the behavior for the default matcher.
func (m *defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
