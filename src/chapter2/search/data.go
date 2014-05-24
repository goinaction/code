package main

import (
	"encoding/json"
	"os"
)

const (
	dataFile = "data/data.json"
)

type (
	// Feed contains information about an rss feed.
	Feed struct {
		Name string `json:"site"`
		Uri  string `json:"link"`
		Type string `json:"type"`
	}
)

// Load retrieves and unmarshals the data for the program.
func Load() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
