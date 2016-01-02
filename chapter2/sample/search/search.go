package search

import (
    "log"
	"sync"
)

// 검색을 처리할 검색기의 매핑 정보를 저장할 맵(map)
var matchers = make(map[string]Matcher)

// 검색 로직을 수행할 Run 함수
func Run(searchTerm string) {
	// 검색할 피드의 목록을 조회한다.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 언버퍼드 채널을 생성하여 화면에 표시할 검색 결과를 전달 받는다.
	results := make(chan *Result)

	// 모든 피드를 처리할 때까지 기다릴 대기 그룹(Wait group)을 설정한다.
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 할
	// 고루틴의 갯수를 설정한다.
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행한다.
	for _, feed := range feeds {
		// 검색을 위해 검색기를 조회한다.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴을 실행힌다.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 모든 작업이 완료되었는지를 모니터링할 고루틴을 실행한다.
	go func() {
		// 모든 작업이 처리될 때까지 기다린다.
		waitGroup.Wait()

		// Display 함수에게 프로그램을 종료할 수 있음을
		// 알리기 위해 채널을 닫는다.
		close(results)
	}()

	// 검색 결과를 화면에 표시하고
	// 마지막 결과를 표시한 뒤 리턴한다.
	Display(results)
}

// 프로그램에서 사용할 검색기를 등록할 함수를 정의한다.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록 완료:", feedType, " 검색기")
	matchers[feedType] = matcher
}
