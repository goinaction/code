// io.Reader와 io.Writer 인터페이스를 이용하여
// curl을 간략하게 재작성한 예제 프로그램
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 초기화 함수
func init() {
	if len(os.Args) != 2 {
		fmt.Println("사용법: ./example2 <url>")
		os.Exit(-1)
	}
}

// 애플리케이션 진입점
func main() {
	// 웹서버로부터 응답을 받는다.
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 본문을 표준 출력으로 복사한다.
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
