// Copyright Information.
//
// This sample program demonstrations how to use an unbuffered
// channel to simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"time"
)

// main is the entry point for all Go programs.
func main() {
    // Create an unbuffered channel.
	ball := make(chan int)

	// Launch two players
	go Player("A", ball)
	go Player("B", ball)

	// Start the lobby
	ball <- 0

	// Give the players time to play.
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("Hits: %d\\n", <-ball)
}

// Player simulates a person playing the game of tennis.
func Player(name string, ball chan int) {
	for {
	    // Wait for the ball to be hit back to us.
		value := <-ball

		// Increment the hit count by one.
		value++
		fmt.Printf("Player %s Hit %d\\n", name, value)

		// Hit the ball back to the opposing player.
		ball <- value
	}
}