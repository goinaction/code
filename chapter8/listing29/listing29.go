// JSON 문자열을 디코딩하는 예제
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 언마샬링이 필요한 문자열을 포함하는 JSON 문서
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// JSON 문자열을 맵에 언마샬링한다.
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println("이름:", c["name"])
	fmt.Println("제목:", c["title"])
	fmt.Println("연락처")
	fmt.Println("집전화:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("휴대폰:", c["contact"].(map[string]interface{})["cell"])
}
