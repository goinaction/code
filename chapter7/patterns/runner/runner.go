// Copyright Information.

// This example is provided with help by Gabriel Aszalos.

// Package runner manages the running and lifetime of
// a process.
package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// runner runs a set of tasks within a given timeout and can be
// shut down on an operating system interrupt.
type runner struct {
	// interrupt channel reports a signal from the
	// operating system.
	interrupt chan os.Signal

	// complete channel reports that processing is done.
	complete chan error

	// timeout reports that time has run out.
	timeout <-chan time.Time

	// tasks holds a set of functions that are executed
	// synchronously in index order.
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

// Add attaches tasks to the runner. A task is a function that
// takes an int ID.
func (r *runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *runner) Start() {
	// We want to receive all interrupt based signals.
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks on a different goroutine.
	go func() {
		r.complete <- r.run()
		log.Println("Finished work.")
	}()

	for {
		select {
		// Signaled when processing is done.
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
func (r *runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
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
	// Signaled when an interrupt event is sent.
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
