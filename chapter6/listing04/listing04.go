// 단일 스레드 환경에서 고루틴이 스케줄러에 의해
// 본할 실행되는 것을 보여주기 위한 예제
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg는 프로그램의 종료를 대기하기 위해 사용한다.
var wg sync.WaitGroup

// 애플리케이션 진입점
func main() {
	// 스케줄러에 하나의 논리 프로세서만 할당한다.
	runtime.GOMAXPROCS(1)

	// 고루틴마다 하나씩, 두 개의 카운트를 추가한다.
	wg.Add(2)

	// 두 개의 고루틴을 생성한다.
	fmt.Println("고루틴을 실행합니다.")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish.
	fmt.Println("대기 중...")
	wg.Wait()

	fmt.Println("프로그램을 종료합니다.")
}

// 소수 중 처음 5000개를 출력하는 함수
func printPrime(prefix string) {
	// 작업이 완료되면 Done 함수를 호출하도록 예약한다.
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("완료: ", prefix)
}
