// Copyright Information.
//
// This sample program demonstrations how to create race
// conditions in our programs. We don't want to do this.
package main

import (
	"fmt"
	"time"
	"runtime"
)

// Counter is a variable incremented by all goroutines.
var Counter int = 0

// main is the entry point for all Go programs.
func main() {
    // Launch two functions as a goroutine.
	go IncCounter(1)
	go IncCounter(2)

    // Give the goroutines time to run.
	time.Sleep(1 * time.Second)
	fmt.Printf("Final Counter: %d\\n", Counter)
}

// IncCounter increments the package level Counter variable.
func IncCounter(id int) {
	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := Counter

		// Yield the processor.
		runtime.Gosched()

		// Increment our local value of Counter.
		value++

		// Store the value back into Counter.
		Counter = value
	}
}
