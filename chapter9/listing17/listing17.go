// 간단한 웹서비스 예제
package main

import (
	"log"
	"net/http"

	"github.com/goinaction/code/chapter9/listing17/handlers"
)

// 애플리케이션 진입점
func main() {
	handlers.Routes()

	log.Println("웹서비스 실행 중: 포트: 4000")
	http.ListenAndServe(":4000", nil)
}
