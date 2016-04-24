// pool 패키지를 이용하여 데이터베이스 연결 풀을
// 생성하고 활용하는 예제
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/goinaction/code/chapter7/patterns/pool"
)

const (
	maxGoroutines   = 25 // 실행할 수 있는 고루틴의 최대 갯수
	pooledResources = 2  // 풀이 관리할 리소스의 갯수
)

// 공유 자원을 표현한 구조체
type dbConnection struct {
	ID int32
}

// dbConnection 타입이 풀에 의해 관리될 수 있도록 
// io.Closer 인터페이스를 구현한다.
// Close 메서드는 자원의 해제를 담당한다.
func (dbConn *dbConnection) Close() error {
	log.Println("닫힘: 데이터베이스 연결", dbConn.ID)
	return nil
}

// 각 데이터베이스에 유일한 id를 할당하기 위한 변수
var idCounter int32

// 풀이 새로운 리소스가 필요할 때 호출할
// 팩토리 메서드
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("생성: 새 데이터베이스 연결", id)

	return &dbConnection{id}, nil
}

// 애플리케이션 진입점
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 데이터베이스 연결을 관리할 풀을 생성한다.
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 풀에서 데이터베이스 연결을 가져와 질의를 실행한다.
	for query := 0; query < maxGoroutines; query++ {
		// 각 고루틴에는 질의 값의 복사본을 전달해야 한다.
		// 그렇지 않으면 고루틴들이 동일한 질의 값을
		// 공유하게 된다.
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// 고루틴의 실행이 종료될 때까지 대기한다.
	wg.Wait()

	// 풀을 닫는다.
	log.Println("프로그램을 종료합니다.")
	p.Close()
}

// 데이터베이스 연결 리소스 풀을 테스트한다.
func performQueries(query int, p *pool.Pool) {
	// 풀에서 데이터베이스 연결 리소스를 획득한다.
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 데이터베이스 연결 리소스를 다시 풀로 되돌린다.
	defer p.Release(conn)

	// 질의문이 실행되는 것처럼 얼마간의 시간 동안 대기한다.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("질의: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
