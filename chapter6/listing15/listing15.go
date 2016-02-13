// atomic 패키지의 함수들을 이용하여
// 숫자 타입의 값을 안전하게 읽고 쓰는
// 방법을 보여주는 예제
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 실행 중인 고루틴들의 종료 신호로 사용될 플래그
	shutdown int64

	// 프로그램이 종료될 때까지 대기할 WaitGroup
	wg sync.WaitGroup
)

// 애플리케이션 진입점
func main() {
	// 고루틴 당 하나씩, 총 두 개의 카운터를 추가한다.
	wg.Add(2)

	// 두 개의 고루틴을 생성한다.
	go doWork("A")
	go doWork("B")

	// 고루틴이 실행될 시간을 할애한다.
	time.Sleep(1 * time.Second)

	// 종료 신호 플래그를 설정한다.
	fmt.Println("프로그램 종료!")
	atomic.StoreInt64(&shutdown, 1)

	// 고루틴의 실행이 종료될 때까지 대기한다.
	wg.Wait()
}

// 필요한 작업을 실행하다가 종료 플래그를 검사하여
// 일찍 종료되는 함수를 흉내내는 함수
func doWork(name string) {
	// 함수 실행이 종료되면 main 함수에 알리기 위해 Done 함수 호출을 에약한다.
	defer wg.Done()

	for {
		fmt.Printf("작업 진행 중: %s\n", name)
		time.Sleep(250 * time.Millisecond)

		// 종료 플래그를 확인하고 실행을 종료한다.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("작업을 종료합니다: %s\n", name)
			break
		}
	}
}
