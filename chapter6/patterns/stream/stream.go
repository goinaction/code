// Copyright Information.
//
// This sample program demonstrations how to use different buffered
// channels to process work.
package main

import (
	"log"
	"sync"

	"github.com/goinaction/code/chapter6/patterns/stream/station"
	"github.com/goinaction/code/chapter6/patterns/stream/timezone"
)

const (
	buffer           = 30
	totalWorkers     = 10
	geoNamesUserName = "ardanstudios"
)

func main() {
	FindTimezones()
}

// FindTimezones use concurrency to retrieve timezones for a set of stations.
func FindTimezones() {
	// This channel buffers all the work coming in off the stream.
	stream := make(chan *station.Station, buffer)

	// This channel contains the current work being processed.
	work := make(chan *station.Station, totalWorkers)

	// This channel is used to communicate work that has been processed
	// back to the main processing routine.
	processed := make(chan *station.Station, 1)

	totalStations := StartStream(stream)
	LaunchWorkRoutines(work, processed)
	stations := ProcessWork(stream, work, processed, totalStations)

	// Shutdown the goroutines that perform the work
	close(work)

	// Display the results for all the processed work.
	for _, station := range stations {
		if station.Err != nil {
			log.Printf("Station[%s]\tERROR[%s]\n", station.Name, station.Err)
		} else {
			log.Printf("Station[%s]\tTZ[%s]\n", station.Name, station.Timezone.TimezoneId)
		}
	}
}

// StartStream simulates the loading of work from a stream into the stream channel so
// it can be processed.
func StartStream(stream chan<- *station.Station) int {
	// Retrieve the set of stations to process.
	stations := station.LoadStations()

	// Launch a goroutine to load these stations
	// into the stream channel.
	go func() {
		for _, station := range stations {
			stream <- station
		}

		// Close the stream channel once the last station
		// is loaded. This will let us determine when the last
		// station is pulled from the stream.
		close(stream)
	}()

	// Return the number of stations we will be processing.
	return len(stations)
}

// LaunchWorkRoutines launch the goroutines that perform the actual work. These goroutines
// call into the GeoNames api to retrieve the timezone information for each station.
func LaunchWorkRoutines(work <-chan *station.Station, processed chan<- *station.Station) {
	for worker := 0; worker <= totalWorkers; worker++ {
		// Launch a goroutine to process work.
		go func() {
			// Pull a station off the work channel.
			for station := range work {
				log.Printf("Work\t: Processing\t: %s\n", station.Name)

				// Call into the geonames api.
				station.Timezone, station.Err = timezone.RetrieveGeoNamesTimezone(
					station.Location.Coordinates[1],
					station.Location.Coordinates[0],
					geoNamesUserName)

				// Push the station on the processed channel
				// so it can be saved and returned.
				processed <- station
			}

			log.Println("Work\t: Shutdown")
		}()
	}
}

// ProcessWork coordinates the work of retrieving timezone information for all the work in the stream. It move
// work from the stream into the work channel and receives processed worked.
func ProcessWork(stream <-chan *station.Station, work chan<- *station.Station, processed <-chan *station.Station, totalStations int) []*station.Station {
	// Create a waitgroup to wait for all the stations to be processed.
	var waitGroup sync.WaitGroup
	waitGroup.Add(totalStations)

	// Slice of processes stations to be returned.
	stations := make([]*station.Station, totalStations)

	go func() {
		streamBuffer := stream // Using a temp vartiable to control the flow of the stream.
		streamClosed := false  // Flag to determine when the stream is closed.
		busyWorkers := 0       // Tracks the number of worker routines that are busy.
		completed := 0         // Tracks the number of stations completed.

		for {
			select {
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
				work <- station

				// If all the work goroutines are busy, don't process any
				// more stations from the stream. Let's get some work done.
				if busyWorkers == totalWorkers {
					log.Printf("Stream\t: Paused\t: Count[%d]\n", busyWorkers)
					streamBuffer = nil
				}

			// Work that has been processed.
			case station := <-processed:
				busyWorkers--
				if station.Err != nil {
					log.Printf("Work\t: ERROR\t: Name[%s] ERROR[%s] Count[%d]\n", station.Name, station.Err, busyWorkers)
				} else {
					log.Printf("Work\t: Completed\t: Name[%s] TZ[%s] Count[%d]\n", station.Name, station.Timezone.TimezoneId, busyWorkers)
				}

				// Store the station in the final slice.
				stations[completed] = station
				completed++

				// Report that work for this station is complete.
				waitGroup.Done()

			// Neither the stream nor the processed channel has work to do. Determine
			// if we open the stream again.
			default:
				if streamBuffer == nil && (busyWorkers < totalWorkers) && !streamClosed {
					log.Printf("Stream\t: Opened\t: Count[%d]\n", busyWorkers)
					streamBuffer = stream
				}
			}
		}
	}()

	// Wait until all the stations have been processed.
	waitGroup.Wait()

	return stations
}
