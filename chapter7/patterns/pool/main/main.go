// Copyright Information.

// This example is provided with help by Gabriel Aszalos.

// This sample program demostrates how to use the pool package
// to share a simulated set of database connections.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/goinaction/code/chapter7/patterns/pool"
)

const (
	maxGoroutines   = 25 // the number of routines to use.
	pooledResources = 2  // number of resources in the pool
)

// dbConnection simulates a resource to share.
type dbConnection struct {
	ID int32
}

// Close implements the interface for the pool package.
// Close performs any resource release management.
func (dbConn *dbConnection) Close() {
	fmt.Println("Close: Connection", dbConn.ID)
}

// isCounter provides support for giving each
// connection a unique id.
var idCounter int32

// createConnection is a factory method called by the pool
// framework when new connections are needed.
func createConnection() (pool.Resource, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

// main is the entry point for all Go programs.
func main() {
	var wg sync.WaitGroup // waits for program to finish
	wg.Add(maxGoroutines)

	// Create the buffered channel to hold
	// and manage the connections.
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		fmt.Println(err)
	}

	// Perform queries using a connection from the pool.
	for query := 0; query < maxGoroutines; query++ {
		// Each goroutine needs its own copy of the query
		// value else they will all be sharing the same query
		// variable.
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Close the pool.
	fmt.Println("*****> Shutdown Program.")
	p.Close()
}

// performQueries tests the resource pool of connections.
func performQueries(query int, p pool.AcquireReleaseCloser) {
	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Release the connection back to the pool.
	defer p.Release(conn)

	// Wait to simulate a query response.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
