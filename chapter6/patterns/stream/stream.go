// Copyright Information.
//
// This sample program demonstrations how to use different buffered
// channels to process work.
package main

import (
	"log"
	"sync"
	"time"

	"github.com/goinaction/code/chapter6/patterns/stream/station"
	"github.com/goinaction/code/chapter6/patterns/stream/timezone"
)

const (
	geoNamesUserName = "ardanstudios"
)

// processor manages a stream of work.
type processor struct {
	// This channel is used to shutdown the program.
	shutdown chan struct{}

	// The WaitGroup keeps track of all running goroutines.
	waitGroup sync.WaitGroup

	// This channel buffers all the work coming in off the stream.
	stream chan *station.Station

	// This channel contains the current work being processed.
	work chan *station.Station

	// This channel is used to communicate work that has been processed
	// back to the main processing routine.
	processed chan *station.Station

	// Total number of workers to process work.
	totalWorkers int
}

// FindTimezones use concurrency to retrieve timezones for a set of stations.
func main() {
	p := processor{
		shutdown:     make(chan struct{}),
		stream:       make(chan *station.Station, 30),
		work:         make(chan *station.Station, 10),
		processed:    make(chan *station.Station, 1),
		totalWorkers: 10,
	}

	p.Run()
	time.Sleep(5 * time.Second)
	p.Shutdown()
}

// Run launches the goroutines to process the stream.
func (p *processor) Run() {
	p.LaunchWorkRoutines()
	p.ProcessWork()
	p.StartStream()
}

// Shutdown shutdown all the running goroutines.
func (p *processor) Shutdown() {
	log.Println("Shutdown\t: Started")

	close(p.shutdown)
	p.waitGroup.Wait()

	log.Println("Shutdown\t: Completed")
}

// StartStream simulates the loading of work from a stream into the stream channel so
// it can be processed.
func (p *processor) StartStream() {
	p.waitGroup.Add(1)

	// Launch a goroutine to load these stations
	// into the stream channel.
	go func() {
		for {
			log.Println("Stream\t: Send Burst")

			// Read the file for work to be accomplished.
			stations := station.LoadStations()

			// Load the work into the stream channel.
			for _, station := range stations {
				p.stream <- station
			}

			select {
			// If we are being signaled to shutdown, do so.
			case <-p.shutdown:
				log.Println("Stream\t: Shutdown")
				p.waitGroup.Done()
				return

			// Sleep for two seconds and send another burst.
			case <-time.After(2 * time.Second):
			}
		}
	}()
}

// LaunchWorkRoutines launch the goroutines that perform the actual work. These goroutines
// call into the GeoNames api to retrieve the timezone information for each station.
func (p *processor) LaunchWorkRoutines() {
	p.waitGroup.Add(p.totalWorkers)

	for worker := 0; worker <= p.totalWorkers; worker++ {
		// Launch a goroutine to process work.
		go func() {
			for {
				select {
				case <-p.shutdown:
					log.Println("Work\t: Shutdown")
					p.waitGroup.Done()
					return

				// Pull a station off the work channel.
				case station := <-p.work:
					log.Printf("Work\t: Processing\t: %s\n", station.Name)

					// Call into the geonames api.
					station.Timezone, station.Err = timezone.RetrieveGeoNamesTimezone(
						station.Location.Coordinates[1],
						station.Location.Coordinates[0],
						geoNamesUserName)

					// Before we write to the channel, check to see if
					// we are shutting down.
					select {
					case <-p.shutdown:
						log.Println("Work\t: Shutdown")
						p.waitGroup.Done()
						return

					default:
					}

					// Push the station on the processed channel
					// so it can be saved and returned.
					p.processed <- station
				}
			}
		}()
	}
}

// ProcessWork coordinates the work of retrieving timezone information for all the work in the stream. It move
// work from the stream into the work channel and receives processed worked.
func (p *processor) ProcessWork() {
	p.waitGroup.Add(1)

	go func() {
		streamBuffer := p.stream // Using a temp vartiable to control the flow of the stream.
		streamClosed := false    // Flag to determine when the stream is closed.
		busyWorkers := 0         // Tracks the number of worker routines that are busy.

		for {
			select {
			case <-p.shutdown:
				log.Println("Processor\t: Shutdown")
				p.waitGroup.Done()
				return

			// Pull work off the stream.
			case station, ok := <-streamBuffer:
				if !ok {
					log.Println("Stream\t: Closed\t:")
					streamClosed = true
					streamBuffer = nil
					continue
				}

				busyWorkers++
				log.Printf("Stream\t: Posting\t: Name[%s] Count[%d]\n", station.Name, busyWorkers)

				// Send the station to the work channel for processing.
				p.work <- station

				// If all the work goroutines are busy, don't process any
				// more stations from the stream. Let's get some work done.
				if busyWorkers == p.totalWorkers {
					log.Printf("Stream\t: Paused\t: Count[%d]\n", busyWorkers)
					streamBuffer = nil
				}

			// Work that has been processed.
			case station := <-p.processed:
				if station == nil {
					log.Println("Process\t: Shutdown")
					p.waitGroup.Done()
					return
				}

				busyWorkers--
				if station.Err != nil {
					log.Printf("Work\t: ERROR\t: Name[%s] ERROR[%s] Count[%d]\n", station.Name, station.Err, busyWorkers)
				} else {
					log.Printf("Work\t: Completed\t: Name[%s] TZ[%s] Count[%d]\n", station.Name, station.Timezone.TimezoneId, busyWorkers)
				}

			// Neither the stream nor the processed channel has work to do. Determine
			// if we open the stream again.
			default:
				if streamBuffer == nil && (busyWorkers < p.totalWorkers) && !streamClosed {
					log.Printf("Stream\t: Opened\t: Count[%d]\n", busyWorkers)
					streamBuffer = p.stream
				}
			}
		}
	}()
}
