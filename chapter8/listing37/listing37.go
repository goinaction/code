// 표준 라이브러리의 각기 다른 위치의 함수들이
// io.Writer 인터페이스를 사용하는 사례를 보여주는 예제
package main

import (
	"bytes"
	"fmt"
	"os"
)

// 애플리케이션 진입점
func main() {
	// Buffer 값을 생성한 후 버퍼에 문자열을 출력한다.
	// 이 때 io.Writer 인터페이스를 구현한 Write 메서드를 호출한다.
	var b bytes.Buffer
	b.Write([]byte("안녕하세요 "))

	// 버퍼에 문자열을 결합하기 위해 Fprintf 함수를 호출한다.
	// 이 때 bytes.Buffer 값의 주소를 io.Writer 타입 매개변수에 전달한다.
	fmt.Fprintf(&b, "Golang!")

	// 버퍼의 콘텐츠를 표준 출력 장치에 쓴다.
	// 이 때 io.Writer 타입의 매개변수에 os.File 값의 주소를 전달한다.
	b.WriteTo(os.Stdout)
}
