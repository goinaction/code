// JSON 문자열을 마샬링하는 방법을 보여주는 예제
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 키/값의 쌍을 가지는 맵을 생성한다.
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 맵을 JSON 문자열로 마샬링한다.
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println(string(data))
}
