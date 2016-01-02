package main

import (
    "log"
    "os"

    _ "github.com/webgenie/go-in-action/chapter2/sample/matchers"
    "github.com/webgenie/go-in-action/chapter2/sample/search"
)

// init 함수는 main 함수보다 먼저 호출된다.
func init() {
    // 표준 출력으로 로그를 출력하도록 변경한다.
    log.SetOutput(os.Stdout)
}

// main 함수는 프로그램의 진입점이다.
func main() {
    // 지정된 검색어로 검색을 수행한다.
    search.Run("Sherlock Holmes")
}
