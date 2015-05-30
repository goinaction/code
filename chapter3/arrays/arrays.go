package main

import "fmt"



func main() {

	//START:declareArrays
	var intArray [10]int64
	//END:declareArrays

	//START:assigning
	intArray[0] = 5
	//END:assigning

	//START:arraylit
	serverNames := [5]string{"hoth","bespin","coruscant","endor","kessel"}
	//END:arraylit

	intArray[1] = 3
	intArray[2] = 2
	intArray[3] = 1
	intArray[4] = 0

	//START:slice
	intSlice := intArray[3:]
	//END:slice

	//START:make
	newSlice := make([]int, 10, 100)
	//END:make

	fmt.Println(newSlice)
	fmt.Println(intArray)
	fmt.Println(serverNames)
	fmt.Println(intSlice)
	
	//<start id="iter"/>
	for i,v := range intArray {
		fmt.Printf("Key %d has value %d", i,v)
	}
	//<end id="iter"/>
	//START:iterUnderscore
	for _,v := range intArray {
		fmt.Printf("Value %d\n", v)
	}
	//END:iterUnderscore

}
