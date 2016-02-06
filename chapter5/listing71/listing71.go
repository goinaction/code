// 노출된 구조체 타입의 비노출 타입 필드에
// 접근이 불가능하다는 것을 설명하기 위한 예제
package main

import (
	"fmt"

	"github.com/webgenie/go-in-action/chapter5/listing71/entities"
)

// 애플리케이션 진입점
func main() {
	// entities 패키지의 User 타입의 값을 생성한다.
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// ./example71.go:16: unknown entities.User field 'email' in
	//                    struct literal

	fmt.Printf("사용자: %v\n", u)
}
