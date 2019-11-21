// Package words provides support for counting words.
package words

import "strings"

// CountWords counts the number of words in the specified
// string and returns the count.
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
