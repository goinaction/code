// 테스트 코드 작성법을 소개 하기 위한 예제 단위 테스트 코드
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// ExampleSendJSON 함수는 기본 예제를 제공한다.
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("에러:", err)
	}

	// fmt 패키지를 이용해서 검사할 결과를 표준 출력 장치에 출력한다.
	fmt.Println(u)
	// Output:
	// {Bill bill@ardanstudios.com}
}
