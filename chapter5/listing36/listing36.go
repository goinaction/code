// Go에서 인터페이스의 사용법을 설명하기 위한 예제
package main

import (
	"fmt"
)

// 알림을 수행하는 동작을 정의하는
// notifier 인터페이스를 선언한다.
type notifier interface {
	notify()
}

// 사용자를 표현하는 user 타입을 선언한다.
type user struct {
	name  string
	email string
}

// 포인터 수신자를 이용하여 notify 메서드를 구현한다.
func (u *user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다.: %s<%s>\n",
		u.name,
		u.email)
}

// 애플리케이션 진입점
func main() {
	// User 타입의 값을 생성한 후 알림을 전송한다.
	u := user{"Bill", "bill@email.com"}

	sendNotification(u)

	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
}

// sendNotification 함수는 notifier 인터페이스를 구현하는 값을 매개변수로 전달받아
// 알림을 보내는 기능을 수행한다.
func sendNotification(n notifier) {
	n.notify()
}
