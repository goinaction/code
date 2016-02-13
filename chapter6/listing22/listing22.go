// 버퍼가 없는 채널을 이용해
// 계주 경기를 묘사하는 예제
package main

import (
	"fmt"
	"sync"
	"time"
)

// 프로그램이 종료될 때까지 대기할 WaitGroup
var wg sync.WaitGroup

// 애플리케이션 진입점
func main() {
	// 버퍼가 없는 채널을 생성한다.
	baton := make(chan int)

	// 마지막 주자를 위해 하나의 카운터를 생성한다.
	wg.Add(1)

	// 첫 번째 주자가 경기를 준비한다.
	go Runner(baton)

	// 경기 시작!
	baton <- 1

	// 경기가 끝날 때까지 기다린다.
	wg.Wait()
}

// 계주의 각 주자를 표현하는 Runner 함수
func Runner(baton chan int) {
	var newRunner int

	// 바톤을 전달받을 때까지 기다린다.
	runner := <-baton

	// 트랙을 달린다.
	fmt.Printf("%d 번째 주자가 바톤을 받아 달리기 시작했습니다.\n", runner)

	// 새로운 주자가 교체지점에서 대기한다.
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("%d 번째 주자가 대기합니다.\n", newRunner)
		go Runner(baton)
	}

	// 트랙을 달린다.
	time.Sleep(100 * time.Millisecond)

	// 경기가 끝났는지 검사한다.
	if runner == 4 {
		fmt.Printf("%d 번째 주자가 도착했습니다. 경기가 끝났습니다. \n", runner)
		wg.Done()
		return
	}

	// 다음 주자에게 바톤을 넘긴다.
	fmt.Printf("%d 번째 주자가 %d 번째 주자에게 바톤을 넘겼습니다.\n",
		runner,
		newRunner)

	baton <- newRunner
}
