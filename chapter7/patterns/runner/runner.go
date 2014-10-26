// Copyright Information.

// This example is provided with help by Gabriel Aszalos.

// Package runner manages the running and lifetime of
// the process.
package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// runner runs a set of tasks on a given timeout and shuts down on os.Interrupt.
type runner struct {
	// interrupt channel will be used to signal the runner to shut down.
	interrupt chan os.Signal

	// complete channel will receive the outcome of the timer.
	complete chan error

	// timeout will signal us after the TTL has run out.
	timeout <-chan time.Time

	// tasks holds a set of runner functions that run based on a given ID.
	tasks []func(int)
}

// New returns a new ready-to-use runner.
func New(d time.Duration) *runner {
	return &runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d * time.Second),
	}
}

// Add attaches tasks to the runner. A task is a function that takes an int ID.
func (r *runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Run runs all tasks.
func (r *runner) Start() {
	// We want to receive all interrupt based signals.
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks.
	go func() {
		r.complete <- r.run(r.tasks...)
		log.Println("Finished work.")
	}()

	for {
		select {
		// Signaled when the tasks are complete.
		case err := <-r.complete:
			if err != nil {
				log.Printf("Exiting with error: %s", err)
			}
			return

		// Signaled when we run out of time.
		case <-r.timeout:
			log.Println("Timeout - Killing Program")
			os.Exit(1)
		}
	}
}

// run executes each registered task.
func (r *runner) run(tasks ...func(int)) error {
	for id, task := range tasks {
		// Check for an interrupt signal for the OS.
		if r.gotInterrupt() {
			return errors.New("Early Shutdown")
		}

		// Execute the registered task.
		task(id)
	}

	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *runner) gotInterrupt() bool {
	select {
	// Signaled when an interrupt event is signaled.
	case <-r.interrupt:
		// Stop receiving any further signals.
		signal.Stop(r.interrupt)
		log.Println("Received interrupt.")
		return true

	// Continue running as normal.
	default:
		return false
	}
}
