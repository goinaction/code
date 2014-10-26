// Copyright Information.

// This example is provided with help by Gabriel Aszalos.

// Package pool manages a user defined set of resources.
// Based on the work by Fatih Arslan with his pool package.
package pool

import (
	"errors"
	"fmt"
	"sync"
)

// AcquireReleaseCloser is behavior that need to be implemented
// to use the pool package.
type AcquireReleaseCloser interface {
	Acquire() (Resource, error)
	Release(Resource)
	Close()
}

// An interface allows us to decouple the pool from its implementation, which is
// a good practice for writing testable and maintainable software.

// Resource must be implemented by types to use the pool.
type Resource interface {
	Close()
}

// Pool manages a set of resources that can be shared safely by multiple goroutines.
type pool struct {
	sync.Mutex
	resources chan Resource
	factory   func() (Resource, error)
	closed    bool
}

// ErrInvalidCapacity is returned when there has been an attempt to create an
// unbuffered pool.
var ErrInvalidCapacity = errors.New("Capacity needs to be greater than zero.")

// New creates a pool from a set of factory functions. A pool provides capacity
// number of resources that can be shared safely by multiple goroutines.
func New(fn func() (Resource, error), capacity uint) (AcquireReleaseCloser, error) {
	if capacity == 0 {
		return nil, ErrInvalidCapacity
	}

	return &pool{
		factory:   fn,
		resources: make(chan Resource, capacity),
	}, nil
}

// Acquire retrieves a resource	from the pool.
func (p *pool) Acquire() (Resource, error) {
	select {
	// Check for a free resource.
	case r, ok := <-p.resources:
		fmt.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, errors.New("Pool has been closed.")
		}
		return r, nil

	// Provide a new resource since there are none available.
	default:
		fmt.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release places a new resource onto the pool.
func (p *pool) Release(r Resource) {
	// Secure this operation with the Close operation.
	p.Lock()
	defer p.Unlock()

	// If the pool is closed, discard the resource.
	if p.closed {
		r.Close()
		return
	}

	select {
	// Attempt to place the new resource on the queue.
	case p.resources <- r:
		fmt.Println("Release:", "In Queue")

	// If the queue is already at capacity we close the resource.
	default:
		fmt.Println("Release:", "Closing")
		r.Close()
	}
}

// Close will shutdown the pool and close all existing resources.
func (p *pool) Close() {
	// Secure this operation with the Release operation.
	p.Lock()
	defer p.Unlock()

	// If the pool is already close, don't do anything.
	if p.closed {
		return
	}

	// Toggle the flag
	p.closed = true

	// Close the channel before we drain the channel of its
	// resources. If we don't do this, we will have a deadlock.
	close(p.resources)

	// Close the resources
	for r := range p.resources {
		r.Close()
	}
}
