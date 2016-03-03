// 기본 단위 테스트를 작성하는 방법을 소개하는 예제
package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 함수는 콘텐츠를 다운로드할 수 있는 HTTP GET 기능을 확인한다.
func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("콘텐츠 다운로드 기능 테스트를 시작.")
	{
		t.Logf("\tURL \"%s\" 호출 시 상태 코드가 \"%d\"인지 확인.",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tHTTP GET 요청을 보냈는지 확인.",
					ballotX, err)
			}
			t.Log("\t\tHTTP GET 요청을 보냈는지 확인.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t상태 코드가 \"%d\" 인지 확인. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\t상태 코드가 \"%d\" 인지 확인. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
