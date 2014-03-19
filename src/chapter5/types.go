package main

import "fmt"

//<start id="declarationsimple">
type MilesPerHour int
//<end id="declarationsimple">

func main(){
	var mph MilesPerHour
	mph = 5
	fmt.Println(mph)
}