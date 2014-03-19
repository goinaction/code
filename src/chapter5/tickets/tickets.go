package main

import (
	"fmt"
	"time"
)

//<start id="structdeclaration">
type SpeedingTicket struct {
	DocumentNumber int
	DriverName     string
	LicenseNumber  string
	IssuingOfficer string
	Location       string
	PostedSpeed    int
	ActualSpeed    int
	DateTime       time.Time
}

//<end id="structdeclaration">

func main() {

	//<start id="longway">
	var tix SpeedingTicket //<co id="createvar" />
	tix.DocumentNumber = 84756
	tix.DriverName = "Speedy Gonzales"
	tix.LicenseNumber = "G234-0598-01287"
	tix.IssuingOfficer = "Officer D. Duck" //<co id="setfields" />
	tix.Location = "Intersection of HWY 54 and HWY 41"
	tix.PostedSpeed = 55
	tix.ActualSpeed = 67
	tix.DateTime = time.Date(2013, time.May, 4, 12, 51, 0, 0, time.UTC)
	//<end id="longway">

	//<start id="shortway">
	tix2 := SpeedingTicket{ //<co id="createandassign" />
		DocumentNumber: 84756,
		DriverName:     "Speedy Gonzales",
		LicenseNumber:  "G234-0598-01287",
		IssuingOfficer: "Officer D. Duck",
		Location:       "Intersection of HWY 54 and HWY 41",
		PostedSpeed:    55,
		ActualSpeed:    67,
		DateTime:       time.Date(2013, time.May, 4, 12, 51, 0, 0, time.UTC), //<co id="comma" />
	}
	//<end id="shortway">

	//<start id="shorterway">
	tix3 := SpeedingTicket{
		84756,
		"Speedy Gonzales",
		"G234-0598-01287",
		"Officer D. Duck",
		"Intersection of HWY 54 and HWY 41",
		55,
		67,
		time.Date(2013, time.May, 4, 12, 51, 0, 0, time.UTC),
	}
	//<end id="shorterway">

	fmt.Println(tix, tix2, tix3)

}
