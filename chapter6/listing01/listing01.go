// 고루틴을 생성하는 방법과 스케줄러의 동작을
// 설명하는 예제
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 애플리케이션 진입점
func main() {
	// 스케줄러가 사용할 하나의 논리 프로세서를 할당한다.
	runtime.GOMAXPROCS(1)

	// wg은 프로그램의 종료를 대기하기 위해 사용한다.
	// 각각의 고루틴마다 하나씩, 총 두 개의 카운트를 추가한다.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("고루틴을 실행합니다.")

	// 익명함수를 선언하고 고루틴을 생성한다.
	go func() {
		// main 함수에게 종료를 알리기 위한 Done 함수 호출을 예약한다.
		defer wg.Done()

		// 알파벳을 세 번 출력한다.
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 익명함수를 선언하고 고루틴을 생성한다.
	go func() {
		// main 함수에게 종료를 알리기 위한 Done 함수 호출을 예약한다.
		defer wg.Done()

		// 알파벳을 세 번 출력한다. 
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 고루틴이 종료될 때까지 대기한다.
	fmt.Println("대기 중...")
	wg.Wait()

	fmt.Println("\n프로그램을 종료합니다.")
}
