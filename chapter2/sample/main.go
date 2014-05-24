package main

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
