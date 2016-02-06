// io.Copy 함수에 bytes.Buffer 타입을 사용하는
// 예제 프로그램
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 애플리케이션 진입점
func main() {
	var b bytes.Buffer

	// 버퍼에 문자열을 기록한다.
	b.Write([]byte("안녕하세요"))

	// Fprintf 함수를 이용하여 버퍼에 문자열을 덧붙인다.
	fmt.Fprintf(&b, "Go 인 액션!")

	// 버퍼의 콘텐츠를 표준 출력 장치에 출력한다.
	io.Copy(os.Stdout, &b)
}
