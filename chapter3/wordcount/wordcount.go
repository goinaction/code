// io 패키지의 활용법을 간략히 소개하기 위한 샘플 프로그램
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/webgenie/go-in-action/chapter3/words"
)

// 애플리케이션 진입점
func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("파일을 열 때 오류가 발생했습니다.", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("총 %d 개의 단어를 발견했습니다. \n", count)
}
