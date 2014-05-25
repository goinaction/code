package main

// Using the blank identifier as explicit package name for matchers
// to force the call to the packages init() function.

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
