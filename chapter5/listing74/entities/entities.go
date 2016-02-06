// 시스템에 정의된 사용자를 표현하는
// 타입을 선언하는 패키지
package entities

// 사용자를 표현하는 User 타입을 선언한다.
type user struct {
	Name  string
	Email string
}

// 관리자를 표현하는 Admin 타입을 선언한다.
type Admin struct {
	user   // 포함된 타입을 비노출 타입으로 선언한다.
	Rights int
}
