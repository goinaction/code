// Copyright Information.

// Example is provided with help by Gabriel Aszalos.

// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	"flag"
	"log"
	"time"

	"github.com/goinaction/code/chapter7/patterns/runner"
)

// flagSec is a command line flag to set the timeout in seconds.
var flagSec = flag.Int("ttl", 3, "timeout in seconds")

// main is the entry point for the program.
func main() {
	// Parse all command line flags.
	flag.Parse()
	log.Println("Starting work.")

	// Create a new timer value for this run.
	r := runner.New(time.Duration(*flagSec))

	// Add the tasks to be run and start running.
	r.Add(createTask(1), createTask(2), createTask(1))
	r.Start()

	log.Println("Process Ended")
}

// createTask returns an example task that sleeps for the given
// duration (in seconds).
func createTask(d time.Duration) func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(d * time.Second)
	}
}
