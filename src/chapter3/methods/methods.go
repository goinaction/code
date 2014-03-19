package main

import "time"
import "fmt"

type LogMessage struct {
	Date    time.Time
	Source  string
	Message string
}

//<start id="method"/>
//<start id="declaration"/>
func (m *LogMessage) TimeAgoInSeconds() int64 {
	//<end id="method"/>
	var seconds int64
	seconds = time.Now().Unix() - m.Date.Unix()
	return seconds
}

//<end id="method"/>

//<start id="func"/>
func TimeAgoInSeconds(m *LogMessage) int64 {
	var seconds int64
	seconds = time.Now().Unix() - m.Date.Unix()
	return seconds
}

//<end id="func"/>

func main() {

	lm := &LogMessage{Date: time.Now(),
		Source:  "GoCron",
		Message: "GoCron failed to create a task"}
	time.Sleep(5000e6)
	fmt.Println("It happended ",
		lm.TimeAgoInSeconds(),
		"seconds ago")

}
