// json 패키지와 NewDecoder 함수를 이용하여 
// JSON 응답을 구조체로 디코딩하는 예제
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// 검색 API의 문서를 매핑하기 위한 gResult 구조체
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// 최상위 문서를 표현하기 위한 gResponse 구조체
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 구글에 검색을 실행한다.
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("에러:", err)
		return
	}
	defer resp.Body.Close()

	// JSON 응답을 구조체로 디코딩한다.
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println(gr)

	// 구조체 타입을 보기좋게 출력할 수 있는
	// JSON 문서로 마샬링한다.
	pretty, err := json.MarshalIndent(gr, "", "    ")
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println(string(pretty))
}
