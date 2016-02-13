// 두 개의 고루틴을 이용해
// 테니스 경기를 모방하는 예제
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 프로그램이 종료될 때까지 대기할 WaitGroup
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 애플리케이션 진입점
func main() {
	// 버퍼가 없는 채널을 생성한다.
	court := make(chan int)

	// 고루틴 당 하나씩, 총 두 개의 카운터를 추가한다.
	wg.Add(2)

	// 두 명의 선수가 등장!
	go player("나달", court)
	go player("죠코비치", court)

	// 경기를 시작한다.
	court <- 1

	// 경기가 끝날때까지 기다린다.
	wg.Wait()
}

// 테니스 선수의 행동을 모방하는 player 함수
func player(name string, court chan int) {
	// 함수의 실행이 종료될 때 Done 함수를 호출하도록 예약한다.
	defer wg.Done()

	for {
		// 공이 되돌아올 때까지 기다린다.
		ball, ok := <-court
		if !ok {
			// 채널이 닫혔으면 승리한 것으로 간주한다.
			fmt.Printf("%s 선수가 승리했습니다.\n", name)
			return
		}

		// 랜덤 값을 이용해 공을 받아치지 못했는지 확인한다.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s 선수가 공을 받아치지 못했습니다.\n", name)

			// 채널을 닫아 현재 선수가 패배했음을 알린다.
			close(court)
			return
		}

		// 선수가 공을 받아 친 횟수를 출력하고 그 값을 증가시킨다. 
		fmt.Printf("%s 선수가 %d 번째 공을 받아쳤습니다.\n", name, ball)
		ball++

		// 공을 상대 선수에게 보낸다.
		court <- ball
	}
}
