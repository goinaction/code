// log 패키지의 기본 기능을 소개하는 예제
package main

import (
	"log"
)

func init() {
	log.SetPrefix("추적: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
    // Println 함수는 표준 로거에 메시지를 출력한다.
	log.Println("메시지")

    // Fatalln 함수는 Println() 함수를 실행한 후 os.Exit(1)을 추가로 호출한다.
	log.Fatalln("치명적 오류 메시지")

    // Panicln 함수는 Println() 함수를 호출한 후 panic() 함수를 추가로 호출한다.
	log.Panicln("패닉 메시지")
}
