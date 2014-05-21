package feeds

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	dataFile = "feeds/feeds.json"
)

type (
	// Site contains information about an rss feed.
	Site struct {
		Name string `json:"site"`
		Uri  string `json:"link"`
	}
)

// Load retrieves and unmarshals the data for the program.
func Load() ([]Site, error) {
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

	// Unmarshal the json document into a slice of sites.
	var sites []Site
	err = json.Unmarshal(data, &sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}
