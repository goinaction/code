// 사용자 정의 로거를 생성하는 예제
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 기타 모든 로그
	Info    *log.Logger // 중요한 정보
	Warning *log.Logger // 경고성 정보
	Error   *log.Logger // 치명적인 오류
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("에러 로그 파일을 열 수 없습니다.", err)
	}

	Trace = log.New(ioutil.Discard,
		"추적: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"정보: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"경고: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"에러: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("일반적인 로그 메시지")
	Info.Println("특별한 정보를 위한 로그 메시지")
	Warning.Println("경고성 로그 메시지")
	Error.Println("에러 로그 메시지")
}
