package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("전송된 값: %d ", i)
	}
	wg.Done()
}

// main 함수는 프로그램의 진입점이다. 
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// 채널에 10개의 정수를 보낸다.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
