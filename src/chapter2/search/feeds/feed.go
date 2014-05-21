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
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(pwd + "/" + dataFile)
	if err != nil {
		return nil, err
	}

	var sites []Site
	err = json.Unmarshal(data, &sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}
