// 내부 및 외부 타입이 동일한 인터페이스를 구현할 때의 문제를
// 설명하기 위한 예제 프로그램
package main

import (
	"fmt"
)

// 알림 동작을 정의하는
// notifier 인터페이스를 선언한다.
type notifier interface {
	notify()
}

// 사용자를 표현하는 user 타입을 선언한다.
type user struct {
	name  string
	email string
}

// user 타입의 값을 통해 호출할 수 있는
// notify 메서드를 구현한다.
func (u *user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n",
		u.name,
		u.email)
}

// 더 많은 권한을 가진 관리자를 표현하는 admin 타입을 선언한다. 
type admin struct {
	user
	level string
}

// admin 타입의 값을 통해 호출할 수 있는
// notify 메서드를 구현한다.
func (a *admin) notify() {
	fmt.Printf("관리자에게 메일을 전송합니다: %s<%s>\n",
		a.name,
		a.email)
}

// 애플리케이션 진입점
func main() {
	// admin 타입의 사용자를 생성한다.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 관리자에게 알림을 전송한다.
	// 이 경우 내부 타입이 구현한 인터페이스는
	// 외부 타입으로 승격되지 않는다.
	sendNotification(&ad)

	// 내부 타입의 메서드에 직접 접근할 수 있다.
	ad.user.notify()

	// 내부 타입의 메서드가 승격되지 않는다.
	ad.notify()
}

// sendNotification 함수는 notifier 인터페이스를 구현하는 값을 매개변수로 전달받아
// 알림을 보내는 기능을 수행한다.
func sendNotification(n notifier) {
	n.notify()
}
