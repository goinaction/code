// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing71/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type User from the entities package.
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// ./example71.go:16: unknown entities.User field 'email' in
	//                    struct literal

	fmt.Printf("User: %v\n", u)
}
