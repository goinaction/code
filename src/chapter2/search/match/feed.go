package match

import (
	"encoding/json"
	"os"
)

const (
	dataFile = "data/data.json"
)

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	Uri  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeed reads and unmarshals the feed data file
// into a native slice for use.
func RetrieveFeed() ([]*Feed, error) {
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
