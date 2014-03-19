// Copyright Information.
//
// This sample program demonstrations how to use an unbuffered
// channel to simulate a relay race between four goroutines.
package main

import (
	"fmt"
	"time"
)

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.
	baton := make(chan int)

	// First runner to his mark.
	go Runner(baton)

	// Start the race.
	baton <- 1

	// Give the runners time to race.
	time.Sleep(500 * time.Millisecond)
}

// Runner simulates a person running in the relay race.
func Runner(baton chan int) {
	var newRunner int

	// Wait to receive the baton.
	runner := <-baton

	// Start running around the track.
	fmt.Printf("Runner %d Running With Baton\\n", runner)

	// New runner to the line.
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\\n", newRunner)
		go Runner(baton)
	}

	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\\n", runner)
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d Exchange With Runner %d\\n", runner, newRunner)
	baton <- newRunner
}

