// 프로그램이 지정된 시간보다 오래 실행 중이면
// 자동으로 종료하기 위해 채널을 활용하는
// 방법을 소개하기 위한 예제
package main

import (
	"log"
	"os"
	"time"

	"github.com/goinaction/code/chapter7/patterns/runner"
)

// 프로그램의 실행 시간
const timeout = 3 * time.Second

// 애플리케이션 진입점
func main() {
	log.Println("작업을 시작합니다.")

	// 실행 시간을 이용해 새로운 작업 실행기를 생성한다.
	r := runner.New(timeout)

	// 수행할 작업을 등록한다.
	r.Add(createTask(), createTask(), createTask())

	// 작업을 실행하고 결과를 처리한다.
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("지정된 작업 시간을 초과했습니다.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("운영체제 인터럽트가 발생했습니다.")
			os.Exit(2)
		}
	}

	log.Println("프로그램을 종료합니다.")
}

// 지정된 시간 동안 아무것도 하지 않고 대기하는
// 예제 작업을 생성하는 함수
func createTask() func(int) {
	return func(id int) {
		log.Printf("프로세서 - 작업 #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
