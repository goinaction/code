// work 패키지의 코드를 이용하여
// 고루틴 풀을 활용하는 방법을 보여주는 예제
package main

import (
	"log"
	"sync"
	"time"

	"github.com/goinaction/code/chapter7/patterns/work"
)

// 화면에 출력할 이름들을 슬라이스로 선언한다.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// 이름을 출력하기 위한 구조체
type namePrinter struct {
	name string
}

// Worker 인터페이스를 구현하기 위해 Task 메서드를 선언한다.
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// 애플리케이션 진입점
func main() {
	// 2개의 고루틴을 위한 작업 풀을 생성한다.
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 이름 슬라이스를 반복한다.
		for _, name := range names {
			// 이름을 지정한다.
			np := namePrinter{
				name: name,
			}

			go func() {
				// 실행할 작업을 등록한다.
				// Run 메서드가 리턴되면 해당 작업이 처리된 것으로 간주한다.
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 작업 풀을 종료하고 이미 등록된 작업들이
	// 종료될 때까지 대기한다.
	p.Shutdown()
}