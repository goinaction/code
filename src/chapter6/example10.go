// Copyright Information.
//
// This sample program demonstrations how to implement a work
// queue using a buffered channel. Multiple goroutines can work
// together to process work from a single queue.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	NUMBER_GOROUTINES           = 4   // Number of goroutines to use.
	WORK_LOAD                   = 10  // Amount of work to process.
	TIME_TO_FINISH_MILLISECONDS = 100 // Time to get work done.
)

// init is called to initialize the package by the
// Go runtime prior to any other code being executed.
func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().Unix())
}

// main is the entry point for all Go programs.
func main() {
	// Set the number of threads in the schedulers queue.
	runtime.GOMAXPROCS(NUMBER_GOROUTINES)

	// Create a buffered channel to queue up work.
	work := make(chan string, WORK_LOAD)

	// Launch goroutines to handle the work.
	for worker := 1; worker <= NUMBER_GOROUTINES; worker++ {
		go Worker(work, worker)
	}

	// Post a bunch of work to get done.
	for post := 1; post <= WORK_LOAD; post++ {
		work <- fmt.Sprintf("Work : %d", post)
	}

	// Give the program time to get work done.
	time.Sleep(time.Duration(TIME_TO_FINISH_MILLISECONDS) * time.Millisecond)
}

// Worker is launched as a goroutine to process work from
// the buffered channel queue.
func Worker(work chan string, worker int) {
	for {
		// Wait for work to be assigned.
		message := <-work

		// Display the message as our work.
		fmt.Printf("Worker: %d : %s\\n", worker, message)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}
