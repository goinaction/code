// Copyright Information.
//
// This sample program demonstrations how to run goroutines
// in parallel.
package main

import (
	"fmt"
	"runtime"
	"time"
)

// main is the entry point for all Go programs.
func main() {
    // Allocate two threads in the schedulers thread pool.
    runtime.GOMAXPROCS(2)

	fmt.Println("Starting Go Routines")

	// Launch an anonymous function as a goroutine.
	go func() {
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

    // Launch a second anonymous function as a goroutine.
	go func() {
		for char := 'A'; char < 'A'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

    // Give the goroutines time to run.
	fmt.Println("Waiting To Finish")
	time.Sleep(1 * time.Second)

	fmt.Println("\\nTerminating Program")
}