// Package counters provides alert counter support.
package counters

// alertCounter is an unexported type that
// contains an integer counter for alerts.
type alertCounter int

// New creates and returns values of the unexported
// type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}
