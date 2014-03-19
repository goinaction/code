package main

import (
	"encoding/json"
	"fmt"
)

func main() {
//<start id="anonymous">
	driver := "Gary Gopher"
	vehicle := "Porsche Boxster S"
	data := struct {
		Driver  string
		Vehicle string
	}{
		driver,
		vehicle,
	}
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
//<end id="anonymous">
}
