// HTTP GET의 모의 호출을 사용하는 예제
// 책에서 사용한 예제와는 다소 다른 부분이 있다.
package listing12

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// feed 변수에는 우리가 기대하는 모의 응답 데이터를 대입한다.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

// mockServer 함수는 GET 요청을 처리할 서버에 대한 포인터를 리턴한다.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 함수는 HTTP GET 요청을 이용해 콘텐츠를 다운로드 한 후
// 해당 콘텐츠를 언마샬링 할 수 있는지 확인한다.
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("콘텐츠 다운로드 기능 테스트를 시작.")
	{
		t.Logf("\tURL \"%s\" 호출 시 상태 코드가 \"%d\"인지 확인.",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tHTTP GET 요청을 보냈는지 확인.",
					ballotX, err)
			}
			t.Log("\t\tHTTP GET 요청을 보냈는지 확인.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\t상태 코드가 \"%d\" 인지 확인. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\t상태 코드가 \"%d\" 인지 확인. %v",
				statusCode, checkMark)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatal("\t\t콘텐츠 언마샬링에 실패했습니다.",
					ballotX, err)
			}
			t.Log("\t\t콘텐츠 언마샬링이 성공했습니다.",
				checkMark)

			if len(d.Channel.Items) == 1 {
				t.Log("\t\t피드에 \"1\" 개의 아이템이 존재하는지 확인.",
					checkMark)
			} else {
				t.Error("\t\t피드에 \"1\" 개의 아이템이 존재하는지 확인.",
					ballotX, len(d.Channel.Items))
			}
		}
	}
}

// Item defines the fields associated with the item tag in
// the buoy RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in
// the buoy RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}
