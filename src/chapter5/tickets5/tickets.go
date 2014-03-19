package main

import (
	"encoding/json"
	"fmt"
	"time"
)


type TicketDisposition int

type SpeedingTicket struct {
	DocumentNumber int `json:"ReferenceNumber"`  //<co id="tags" />
	*Driver
	IssuingOfficer string
	Location       string
	PostedSpeed    int
	ActualSpeed    int
	DateTime       time.Time
	disposition    TicketDisposition
	Fine           int
}

type Driver struct {
	LicenseNumber     string
	DriverName        string
	Address           string
	City              string
	State             string
	PostalCode        int
	LicenseExpiration time.Time
}

const ( //<co id="enum" />
	NotGuilty TicketDisposition = iota
	Guilty
	NoContest
	Dismissed
	Paid
)

func (ticket *SpeedingTicket) ChangeDisposition(d TicketDisposition) {
	ticket.disposition = d //<co id="method" />
}

func (ticket *SpeedingTicket) PayInFull(payment int) {
	if payment >= ticket.Fine {
		ticket.disposition = Paid
	}

}

func (ticket *SpeedingTicket) ToJson() ([]byte, error) {
	b, err := json.Marshal(ticket) //<co id="marshal" />
	return b, err
}


func main() {

	var ticketJson = []byte(
		`{"ReferenceNumber":84756,
		"LicenseNumber":"G234-0598-01287",
		"DriverName":"Speedy Gonzales",
		"Address":"123 Any Street",
		"City":"Albuquerque",
		"State":"NM",
		"PostalCode":51234,
		"LicenseExpiration":"2017-05-30T00:00:00Z",
		"IssuingOfficer":"Officer D. Duck",
		"Location":"Intersection of HWY 54 and HWY 41",
		"PostedSpeed":55,
		"ActualSpeed":67,
		"DateTime":"2013-05-04T12:51:00Z",
		"Fine":250}`)

	//<start id="fromjson">
	var ticket SpeedingTicket
	//<end id="fromjson">
	
	//<start id="unmarshal">
	err := json.Unmarshal(ticketJson, &ticket)
	//<end id="unmarshal">

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", ticket)
}
