// 가브리엘 애자로스(Gabriel Aszalos)가 도움을 준 예제
// runner 패키지는 프로세스의 실행 및 수명주기를 관리한다.
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 타입은 주어진 타임아웃 시간 동안 일련의 작업을 수행한다.
// 그리고 운영체제 인터럽트에 의해 실행이 종료된다.
type Runner struct {
	// 운영체제로부터 전달되는 인터럽트 신호를
	// 수신하기 위한 채널
	interrupt chan os.Signal

	// 처리가 종료되었음을 알리기 위한 채널
	complete chan error

	// 지정된 시간이 초과했음을 알리기 위한 채널
	timeout <-chan time.Time

	// 인덱스 순서로 처리될 작업의 목록을
	// 저장하기 위한 슬라이스
	tasks []func(int)
}

// timeout 채널에서 값을 수신하면 ErrTimeout을 리턴한다.
var ErrTimeout = errors.New("시간을 초과했습니다.")

// 운영체제 이벤트를 수신하면 ErrInterrupt를 리턴한다.
var ErrInterrupt = errors.New("운영체제 인터럽트 신호를 수신했습니다.")

// 실행할 Runner 타입 값을 리턴하는 함수
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Runner 타입에 작업을 추가하는 메서드
// 작업은 int형 ID를 매개변수로 전달받는 함수이다.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 저장된 모든 작업을 실행하고 채널 이벤트를 관찰한다.
func (r *Runner) Start() error {
	// 모든 종류의 인터럽트 신호를 수신한다.
	signal.Notify(r.interrupt, os.Interrupt)

	// 각각의 작업을 각기 다른 고루틴을 통해 실행한다.
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 작업 완료 신호를 수신한 경우
	case err := <-r.complete:
		return err

	// 작업 시간 초과 신호를 수신한 경우
	case <-r.timeout:
		return ErrTimeout
	}
}

// 개별 작업을 실행하는 메서드
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// OS로부터 인터럽트 신호를 수신했는지 확인한다.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 작업을 실행한다.
		task(id)
	}

	return nil
}

// 인터럽트 신호가 수신되었는지 확인하는 메서드
func (r *Runner) gotInterrupt() bool {
	select {
	// 인터럽트 이벤트가 발생한 경우
	case <-r.interrupt:
		// 이후에 발생하는 인터럽트 신호를 더 이상 수신하지 않도록 한다.
		signal.Stop(r.interrupt)
		return true

	// 작업을 계속해서 실행하게 한다.
	default:
		return false
	}
}
