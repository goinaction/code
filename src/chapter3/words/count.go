package words

import "strings"

func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
