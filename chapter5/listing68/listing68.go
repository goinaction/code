// 다른 패키지에서 비노출 식별자에 대한 접근이 가능하다는 것을
// 보여주기 위한 예제 프로그램
package main

import (
	"fmt"

	"github.com/webgenie/go-in-action/chapter5/listing68/counters"
)

// 애플리케이션 진입점
func main() {
    // counters 패키지가 노출한 New 함수를 이용하여
    // 비노출 타입의 변수에 접근할 수 있게 된다.
	counter := counters.New(10)

	fmt.Printf("카운터: %d\n", counter)
}
