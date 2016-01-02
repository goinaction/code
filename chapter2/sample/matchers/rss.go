package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/webgenie/go-in-action/chapter2/sample/search"
)

type (
	// item 구조체는 RSS 문서 내의 item 태그에
	// 정의된 필드들에 대응하는 필드들을 선언한다.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image 구조체는 RSS 문서 내의 image 태그에
	// 정의된 필드들에 대응하는 필드들을 선언한다.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 구조체는 RSS 문서 내의 channel 태그에
	// 정의된 필드들에 대응하는 필드들을 선언한다.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 구조체는 RSS 문서에 정의된 필드들에 대응하는 필드들을 정의한다.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// Matcher 인터페이스를 구현하는 rssMatcher 타입을 선언한다.
type rssMatcher struct{}

// init 함수를 통해 프로그램에 검색기를 등록한다.
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search 함수는 지정된 문서에서 검색어를 검색한다.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("피드 종류[%s] 사이트[%s] 주소[%s]에서 검색을 수행합니다.\n", feed.Type, feed.Name, feed.URI)

	// 검색할 데이터를 조회한다.
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// 제목에서 검색어를 검색한다.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장한다.
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 상세 내용에서 검색어를 검색한다.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장한다.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// HTTP Get 요청을 수행해서 RSS 피드를 요청한 후 결과를 디코딩한다.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("검색할 RSS 피드가 정의되지 않았습니다.")
	}

	// 웹에서 RSS 문서를 조회한다.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 함수가 리턴될 때 응답 스트림을 닫는다.
	defer resp.Body.Close()

	// 상태 코드가 200인지를 검사해서
	// 올바른 응답을 수신했는지를 확인한다.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP 응답 오류: %d\n", resp.StatusCode)
	}

	// RSS 피드 문서를 구조체 타입으로 디코드한다.
	// 호출 함수가 에러를 판단할 것이기 때문에 이 함수에서는 에러를 처리하지 않는다.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
