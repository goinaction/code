// 노출된 구조체 타입의 비노출 필드에 접근이 불가능하다는 것을
// 설명하기 위한 예제
package main

import (
	"fmt"

	"github.com/webgenie/go-in-action/chapter5/listing74/entities"
)

// 애플리케이션 진입점
func main() {
	// entities 패키지의 Admin 타입의 값을 생성한다.
	a := entities.Admin{
		Rights: 10,
	}

	// 비노출 타입인 내부 타입의 노출 타입의 필드들에
	// 값을 대입한다.
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("사용자: %v\n", a)
}
