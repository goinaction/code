package station

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"test/book/timezone"
)

const (
	fixture = "station/stations.json"
)

type (

	// StationLocation contains the buoys latitude and longitude location.
	StationLocation struct {
		Type        string    `bson:"type"`
		Coordinates []float64 `bson:"coordinates"`
	}

	// Station contains information for an individual station.
	Station struct {
		Name     string          `bson:"name"`
		LocDesc  string          `bson:"location_desc"`
		Location StationLocation `bson:"location"`
		Timezone *timezone.GeoNamesTimezone
		Err      error
	}
)

// LoadStations retrieves and unmarshals the data for the program.
func LoadStations() []Station {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadFile(pwd + "/" + fixture)
	if err != nil {
		log.Fatalln(err)
	}

	var stations []Station
	err = json.Unmarshal(data, &stations)
	if err != nil {
		log.Fatalln(err)
	}

	return stations
}
