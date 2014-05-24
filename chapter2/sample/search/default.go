package search

// defaultMatcher implements the defaut matcher.
type defaultMatcher struct{}

// init registered the matcher with the program
func init() {
	Register("default", &defaultMatcher{})
}

// Search implements the behavior for the default matcher.
func (m *defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
