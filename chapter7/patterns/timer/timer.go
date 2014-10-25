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

var (
	flagSec = flag.Int("ttl", 3, "timeout in seconds")

	// complete is used to report processing is done.
	complete = make(chan error)

	// shutdown provides system wide notification.
	shutdown = make(chan bool)
)

// main is the entry point for all Go programs.
func main() {
	flag.Parse()

	// Launch the process.
	log.Println("Launching Processors")

	go processor(complete)
	controlLoop()

	// Program finished.
	log.Println("Process Ended")
}

func controlLoop() {
	// We want to receive all interrupt based signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Set the timeout in seconds based on the given flag
	timeout := time.After(time.Duration(*flagSec) * time.Second)
	for {
		select {
		// Interrupt event signaled by the operation system.
		case <-sigChan:
			log.Println("OS INTERRUPT")
			// Signal shutdown
			shutdown <- true
			// By setting the channel to 'nil' we will discard all future sends.
			sigChan = nil

		// We have taken too much time. Kill the app.
		case <-timeout:
			log.Println("Timeout - Killing Program")
			os.Exit(1)

		case err := <-complete:
			// Everything completed within the time given.
			log.Printf("Task Completed: Error[%s]", err)
			return
		}
	}
}

// checkShutdown checks the shutdown flag to determine
// if we have been asked to interrupt processing.
func checkShutdown() bool {
	select {
	case <-shutdown:
		// We have been asked to shutdown cleanly.
		log.Println("checkShutdown - Shutdown Early")
		return true

	default:
		// If the shutdown channel was not closed,
		// presume with normal processing.
	}

	return false
}

// processor provides the main program logic for the program.
func processor(complete chan<- error) {
	log.Println("Processor - Starting")

	defer func() {
		// Capture any potential panic.
		if r := recover(); r != nil {
			log.Println("Processor - Panic", r)
		}

	}()

	// Perform the work and send the returned error back
	complete <- doWork()

	log.Println("Processor - Completed")
}

// doWork simulates task work.
func doWork() error {
	log.Println("Processor - Task 1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}
