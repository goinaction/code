// atomic 패키지의 함수들을 이용하여
// 숫자 타입에 안전하게 접근하는 예제
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// 공유 자원으로 활용될 변수
	counter int64

	// 프로그램이 종료될 때까지 대기할 WaitGroup
	wg sync.WaitGroup
)

// 애플리케이션 진입점
func main() {
	// 고루틴 당 하나씩, 총 두 개의 카운터를 추가한다.
	wg.Add(2)

	// 두 개의 고루틴을 생성한다.
	go incCounter(1)
	go incCounter(2)

	// 고루틴의 실행이 종료될 때까지 대기한다.
	wg.Wait()

	// 최종 결과를 출력한다.
	fmt.Println("최종 결과:", counter)
}

// 패키지 수준에 정의된 counter 변수의 값을 증가시키는 함수
func incCounter(id int) {
	// 함수 실행이 종료되면 main 함수에 알리기 위해 Done 함수 호출을 에약한다.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// counter 변수에 안전하게 1을 더한다.
		atomic.AddInt64(&counter, 1)

		// 스레드를 양보하고 실행 큐로 되돌아간다.
		runtime.Gosched()
	}
}
