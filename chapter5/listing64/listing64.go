// 다른 패키지에서 비노출 식별자에 대한 접근이 차단되는 것을
// 보여주기 위한 예제 프로그램
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing64/counters"
)

// 애플리케이션 진입점 
func main() {
    // 비노출 타입의 변수를 생성하고
    // 10이라는 값으로 초기화한다.
	counter := counters.alertCounter(10)

	// ./listing64.go:15: cannot refer to unexported name
	//                                         counters.alertCounter
	// ./listing64.go:15: undefined: counters.alertCounter

	fmt.Printf("카운터: %d\n", counter)
}
