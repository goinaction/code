// 단어의 개수를 카운트하는 words 패키지
package words

import "strings"

// CountWords 함수는 지정된 문자열에서
// 단어의 개수를 세어 리턴한다.
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
