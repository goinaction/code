// Copyright Information.
//
// This sample program demonstrations how to use a mutex
// to define critical sections of code that need synchronous
// access.
package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

// Counter is a variable incremented by all goroutines.
var Counter int = 0

// Mutex is used to define a critical section of code.
var Mutex sync.Mutex

// main is the entry point for all Go programs.
func main() {
    // Launch two functions as a goroutine.
	go IncCounter(1)
	go IncCounter(2)

    // Give the goroutines time to run.
	time.Sleep(1 * time.Second)
	fmt.Printf("Final Counter: %d\\n", Counter)
}

// IncCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.
func IncCounter(id int) {
	for count := 0; count < 2; count++ {
	    // Only allow one goroutine through this
	    // critical section at a time.
	    Mutex.Lock()

		// Capture the value of Counter.
		value := Counter

		// Yield the processor.
		runtime.Gosched()

		// Increment our local value of counter.
		value++

		// Store the value back into counter.
		Counter = value

		// Release the lock and allow any
		// waiting goroutine through.
		Mutex.Unlock()
	}
}
