// 타입 포함을 이용해 다른 타입을 포함하는 방법과
// 이 경우 내부 및 외부 타입의 관계를 확인하기 위한 예제 프로그램
package main

import (
	"fmt"
)

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
	user  // 포함된 타입
	level string
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

	// 내부 타입의 메서드를 직접 호출할 수 있다.
	ad.user.notify()

	// 내부 타입의 메서드가 승격되었다.
	ad.notify()
}
