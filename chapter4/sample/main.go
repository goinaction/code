package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Relative path to the data file.
const dataFile = "data/latest_obs.txt"

// The list of field names in their file position.
var fieldNames = [...]string{
	"#STN", "LAT", "LON", "YYYY", "MM", "DD", "hh",
	"mm", "WDIR", "WSPD", "GST", "WVHT", "DPD", "APD",
	"MWD", "PRES", "PTDY", "ATMP", "WTMP", "DEWP",
	"VIS", "TIDE"}

// main is the entry point for the program.
func main() {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Schedule the file to be closed.
	defer file.Close()

	// Contains all the field/value mappings for every line.
	var fvMappings []map[string]string

	// Create a reader for the file.
	r := bufio.NewReader(file)
	for {
		// Read all the bytes up to the end of line marker.
		line, err := r.ReadSlice('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		// Capture the field/value mappings for this line.
		fv := make(map[string]string)

		var start int
		var field int
		for index := 0; index < len(line); index++ {
			// If we don't find a space or EOL, check the next byte.
			if line[index] != ' ' && line[index] != '\n' {
				continue
			}

			// If the start and index values are the same, we have more than
			// one space separating the next value.
			if start == index {
				start = index + 1
				continue
			}

			// Slice the value from the line and add the value to the map
			// for the specified field name.
			fv[fieldNames[field]] = string(line[start:index])
			field++
			start = index + 1
		}

		// Append the field/value map to the master collection.
		fvMappings = append(fvMappings, fv)
	}

	// Display all of the field/value maps.
	for _, fv := range fvMappings {
		fmt.Printf("%#v\n\n", fv)
	}
}
