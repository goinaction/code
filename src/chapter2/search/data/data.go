package data

import (
	"encoding/json"
	"io/ioutil"
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
func Load() ([]Feed, error) {
	// Get the current directory we are running inside.
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Read the entire file.
	data, err := ioutil.ReadFile(pwd + "/" + dataFile)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a slice of feeds.
	var feeds []Feed
	err = json.Unmarshal(data, &feeds)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
