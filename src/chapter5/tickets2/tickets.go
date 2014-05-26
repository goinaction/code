package main

import (
	"fmt"
	"time"
)

//<start id="refactoredticket">
type SpeedingTicket struct {
	DocumentNumber int
	Driver         Driver
	IssuingOfficer string
	Location       string
	PostedSpeed    int
	ActualSpeed    int
	DateTime       time.Time
}

//<end id="refactoredticket">

//<start id="drivertype">
type Driver struct {
	LicenseNumber     string
	DriverName        string
	Address           string
	City              string
	State             string
	PostalCode        int
	LicenseExpiration time.Time
}

//<end id="drivertype">

func main() {

	//<start id="longway">
	var tix SpeedingTicket //<co id="createvar" />
	tix.DocumentNumber = 84756

	var d Driver //<co id="createdriver" />
	d.LicenseNumber = "G234-0598-01287"
	d.DriverName = "Speedy Gonzales"
	d.Address = "123 Any Street"
	d.City = "Albuquerque"
	d.State = "NM"
	d.PostalCode = 51234
	d.LicenseExpiration = time.Date(2017, time.May, 30, 0, 0, 0, 0, time.UTC)

	tix.Driver = d //<co id="assigndriver" />

	// 	tix.LicenseNumber = "G234-0598-01287"
	tix.IssuingOfficer = "Officer D. Duck"
	tix.Location = "Intersection of HWY 54 and HWY 41"
	tix.PostedSpeed = 55
	tix.ActualSpeed = 67
	tix.DateTime = time.Date(2013, time.May, 4, 12, 51, 0, 0, time.UTC)
	//<end id="longway">

	//<start id="shorterway">
	tix3 := SpeedingTicket{
		84756,
		Driver{"G234-0598-01287", //<co id="embeddedstructliteral" />
			"Speedy Gonzales",
			"123 Any Street",
			"Albuquerque",
			"NM",
			51234,
			time.Date(2017, time.May, 30, 0, 0, 0, 0, time.UTC),
		}, //<co id="anothercomma" />
		"Officer D. Duck",
		"Intersection of HWY 54 and HWY 41",
		55,
		67,
		time.Date(2013, time.May, 4, 12, 51, 0, 0, time.UTC),
	}
	//<end id="shorterway">

	fmt.Println(tix3)

	//<start id="embeddedselector">
	fmt.Println(tix.Driver.LicenseNumber)

	//<end id="embeddedselector">
}
