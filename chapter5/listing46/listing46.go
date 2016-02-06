// 값의 주소를 알아낼 수 없는 경우를
// 설명하기 위한 예제 프로그램
package main

import "fmt"

// int 타입을 기반 타입으로 duration 타입을 선언한다.
type duration int

// duration 값을 예쁘게 출력하는 함수
func (d *duration) pretty() string {
	return fmt.Sprintf("기간: %d", *d)
}

// 애플리케이션 진입점
func main() {
	duration(42).pretty()

	// ./listing46.go:17: cannot call pointer method on duration(42)
	// ./listing46.go:17: cannot take the address of duration(42)
}
