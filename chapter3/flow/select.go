package main

import (
	"fmt"
)


func sendIntegers(integerStream chan int, done chan bool) {

	for x := 1; x < 100; x++ {
		select {
		case _ = <-done:
			break
		default:
			fmt.Println("no communication\n")
		}
		integerStream <- x
	}
}

func catchIntegers(ints chan int, done chan bool){
	var caught int
	select{
	case ints <- caught :
		fmt.Printf("Received %d", caught)
		if caught > 89 {
			done <- true
		}
	}
}

func kickOff(){
	var integers chan int
	var finished chan bool
	
	go sendIntegers(integers, finished)
	go catchIntegers(integers, finished)
	
}