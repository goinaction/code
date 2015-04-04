// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing62/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	counter := counters.alertCounter(10)

	// ./listing62.go:15: cannot refer to unexported name
	//                                         counters.alertCounter
	// ./listing62.go:15: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
