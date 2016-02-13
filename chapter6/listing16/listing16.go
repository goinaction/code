// 접근 동기화가 필요한 코드에
// 뮤텍스를 이용해 임계 지역을 생성해서
// 경합 상태를 해결하는 예제
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 공유 자원으로 활용될 변수
	counter int

	// 프로그램이 종료될 때까지 대기할 WaitGroup
	wg sync.WaitGroup

	// 코드의 임계 지역을 설졍할 때 사용할 뮤텍스
	mutex sync.Mutex
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
	fmt.Printf("최종 결과: %d\n", counter)
}

// 패키지 수준에 정의된 counter 변수의 값을 
// 뮤텍스를 이용해 안전하게 증가시키는 함수
func incCounter(id int) {
	// 함수 실행이 종료되면 main 함수에 알리기 위해 Done 함수 호출을 에약한다.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 이 임계 지역에는 한 번에 하나의
		// 고루틴만이 접근할 수 있다.
		mutex.Lock()
		{
			// counter 변수의 값을 읽는다.
			value := counter

			// 스레드를 양보하여 큐로 돌아가도록 한다.
			runtime.Gosched()

			// 현재 카운터 값을 증가시킨다.
			value++

			// 원래 변수에 증가된 값을 다시 저장한다.
			counter = value
		}
		mutex.Unlock()
		// 대기 중인 다른 고루틴이 접근할 수 있도록
		// 잠금을 해제한다.
	}
}
