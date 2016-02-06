// Go에서 메서드를 선언하는 방법과 컴파일러가 함수를 어떻게 지원하는지를
// 설명하기 위한 예제 프로그램
package main

import (
	"fmt"
)

// 프로그램의 사용자를 표현하는 user 타입
type user struct {
	name  string
	email string
}

// 값 수신자와 함께 notify 메서드를 선언한다.
func (u user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n",
		u.name,
		u.email)
}

// 포인터 수신자와 함께 changeEmail 메서드를 선언한다.
func (u *user) changeEmail(email string) {
	u.email = email
}

// 애플리케이션 진입점
func main() {
	// user 타입의 값을 이용하여 값 수신자에 선언한
	// 메서드를 호출한다.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// user 타입의 포인터를 이용하여 값 수신자에 선언한
	// 메서드를 호출한다.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user 타입의 값을 이용하여 포인터 수신자에 선언한
	// 메서드를 호출한다.
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// user 타입의 포인터를 이용하여 포인터 수신자에 선언한
	// 메서드를 호출한다.
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
