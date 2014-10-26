// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"
)

// flagSec is a command line flag to set the timeout in seconds
var flagSec = flag.Int("ttl", 3, "timeout in seconds")

// timer runs a set of workers on a given timeout and shuts down on os.Interrupt
type timer struct {
	// the interrupt channel will be used to signal the runner to shut down
	interrupt chan os.Signal
	// the complete channel will receive the outcome of the timer
	complete chan error
	// timeout will signal us after the TTL has run out
	timeout <-chan time.Time
	// workers holds a set of worker functions that run based on a given ID
	workers []func(int)
}

// NewTimer returns a new ready-to-use timer.
func NewTimer(d time.Duration) *timer {
	t := &timer{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d * time.Second),
		workers:   make([]func(int), 0, 5),
	}
	// We want to receive all interrupt based signals.
	signal.Notify(t.interrupt, os.Interrupt)
	return t
}

// Add attaches workers to the timer. A worker is a function that takes an int ID.
func (t *timer) Add(workers ...func(int)) {
	t.workers = append(t.workers, workers...)
}

// Run runs all workers.
func (t *timer) Start() {
	// Run work async
	go func() {
		t.complete <- t.run(t.workers...)
		log.Println("Finished work.")
	}()

	for {
		select {
		// Task completed
		case err := <-t.complete:
			if err != nil {
				log.Printf("Exiting with error: %s", err)
			}
			return

		// We have taken too much time. Kill the app.
		case <-t.timeout:
			log.Println("Timeout - Killing Program")
			os.Exit(1)
		}
	}
}

// doWork simulates task work.
func (t *timer) run(workers ...func(int)) error {
	for id, wrk := range workers {
		if !t.canContinue() {
			return errors.New("Early Shutdown")
		}
		wrk(id)
	}
	return nil
}

// canContinue verifies if the interrupt signal has been sent
func (t *timer) canContinue() bool {
	select {
	// check if we are being signaled to shut down
	case <-t.interrupt:
		log.Println("Received interrupt.")
		return false
	// otherwise continue as normal
	default:
		return true
	}
}

// sleeper returns a worker that sleeps for the given duration (in seconds)
func sleeper(d time.Duration) func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(d * time.Second)
	}
}

func main() {
	// Parse all command line flags.
	flag.Parse()
	log.Println("Starting work.")

	timer := NewTimer(time.Duration(*flagSec))
	timer.Add(sleeper(1), sleeper(2), sleeper(1))
	timer.Start()
	// Program finished.
	log.Println("Process Ended")
}
