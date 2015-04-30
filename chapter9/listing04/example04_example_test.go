// Sample example test for the LogResponse function.
package main_test

import (
	"log"
	"os"

	ex "github.com/goinaction/code/chapter9/listing04"
)

// ExampleLogResponse provides a basic example test example.
func ExampleLogResponse() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	ex.LogResponse(&u)
	// Output:
	// {
	//     "Name": "Bill",
	//     "Email": "bill@ardanstudios.com"
	// }
}
