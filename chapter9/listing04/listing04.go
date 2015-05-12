// This sample code implement a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/goinaction/code/chapter9/listing04/handlers"
)

// main is the entry point for the application.
func main() {
	Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}

// Routes sets the routes for the web service.
func Routes() {
	http.HandleFunc("/sendjson", handlers.SendJSON)
}
