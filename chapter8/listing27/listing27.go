// JSON 문자열을 디코딩하는 예제
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON 문자열을 표현하는 Contact 구조체
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// 문자열을 가지고 있는 JSON 문서
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// JSON 문자열을 변수에 언마샬링한다.
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println(c)
}
