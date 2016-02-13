// 버퍼가 있는 채널을 이용해
// 미리 정해진 고루틴의 갯수만큼
// 다중 작업을 수행하는 예제
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 실행할 고루틴의 갯수
	taskLoad         = 10 // 처리할 작업의 갯수
)

// 프로그램이 종료될 때까지 대기할 WaitGroup
var wg sync.WaitGroup

// Go 런타임이 다른 코드를 실행하기에 앞서
// 패키지의 초기화를 위해 호출하는 함수
func init() {
	// 랜덤 값 생성기를 초기화한다.
	rand.Seed(time.Now().Unix())
}

// 애플리케이션 진입점
func main() {
	// 작업 부하를 관리하기 위한 버퍼가 있는 채널을 생성한다.
	tasks := make(chan string, taskLoad)

	// 작업을 처리할 고루틴을 실행한다.
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 실행할 작업을 추가한다.
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업: %d", post)
	}

	// 작업을 모두 처리하면
	// 채널을 닫는다.
	close(tasks)

	// 모든 작업이 처리될 때까지 대기한다.
	wg.Wait()
}

// 버퍼가 있는 채널에서 수행할 작업을
// 가져가는 고루틴
func worker(tasks chan string, worker int) {
	// 함수가 리턴될 때 Done 함수를 호출하도록 예약한다.
	defer wg.Done()

	for {
		// 작업이 할당될 때까지 대기한다.
		task, ok := <-tasks
		if !ok {
			// 채널이 닫힌 경우
			fmt.Printf("작업자: %d : 종료합니다.\n", worker)
			return
		}

		// 작업을 시작하는 메시지를 출력한다.
		fmt.Printf("작업자: %d : 작업 시작: %s\n", worker, task)

		// 작업을 처리하는 것을 흉내내기 위해 임의의 시간동안 대기한다.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 작업이 완료되었다는 메시지를 출력한다.
		fmt.Printf("작업자: %d : 작업 완료: %s\n", worker, task)
	}
}
