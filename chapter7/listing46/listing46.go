// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// main is the entry point for the application.
func main() {
	// r here is a response, and r.Body is an io.Reader.
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// Create a file to persist the response.
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Use MultiWriter so we can write to stdout and
	// a file on the same write operation.
	dest := io.MultiWriter(os.Stdout, file)

	// Read the response and write to both destinations.
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
