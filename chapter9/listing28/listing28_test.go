// 정수를 문자열로 변환하는 가장 빠른 방법을 테스트하기 위해 벤치마킹을 활용하는 예제
// 첫 번째 방법은 fmt.Sprintf 함수를, 두 번째 방법은 strconv.FormatInt 함수를 테스트하며,
// 마지막으로는 strconv.Itoa 함수를 테스트한다.
package listing05_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf 함수는 fmt.Sprintf 함수의
// 성능을 테스트한다.
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat 함수는 strconv.FormatInt 함수의
// 성능을 테스트한다.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 함수는 strconv.Itoa 함수의
// 성능을 테스트한다.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
