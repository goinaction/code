// Copyright Information.

// Example is provided with help by Jason Waldrip.

// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/goinaction/code/chapter7/patterns/work"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m *namePrinter) Work() {
	fmt.Println(m.name)
	time.Sleep(time.Second)
}

// main is the entry point for all Go programs.
func main() {
	// Create a work value with 2 goroutines.
	w := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When RunTask
				// returns we know it is being handled.
				w.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shutdown the work and wait for all existing work
	// to be completed.
	w.Shutdown()
}
