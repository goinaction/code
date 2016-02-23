// 제이슨 월드립(Json Waldrip)이 도움을 준 예제
// 특정 작업을 수행하는 고루틴의 풀을 관리하는 패키지
package work

import "sync"

// work 풀을 사용하려는 타입은
// Worker 인터페이스를 구현해야 한다.
type Worker interface {
	Task()
}

// Worker 인터페이스를 실행하는 고루틴의
// 풀을 제공하기 위한 Pool 구조체
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// 새로운 work 풀을 생성하는 함수
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// 풀에 새로운 작업을 추가하는 메서드
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// 모든 고루틴을 종료할 때까지 대기하는 메서드
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
