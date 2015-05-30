package main

import (
	"fmt"
	"log"
	"os"
)

	//<start id="stackeddefer"/>
func showDefer(){
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	return
}
	//<end id="stackeddefer"/>

func getX() int {
	return 5
}

func main() {

	//<start id="defer"/>
	var logger *log.Logger

	f, err := os.Create("flow.log") //<co id="openfile" />

	defer f.Close() //<co id="deferclose" />
	if err != nil {
		fmt.Println(err)
	}
	logger = log.New(f, "", log.Ldate|log.Ltime)
	//<end id="defer"/>

	//<start id="if"/>
	if 1 > 2 {
		fmt.Println("Logic Error.")
	}
	//<end id="if"/>

	//<start id="ifpre"/>
	if x := getX(); x > 2 {
		fmt.Println("X is greater than 2")
	}
	//<end id="ifpre"/>

	//<start id="else"/>
	if 1 > 2 {
		fmt.Println("Logic Error.")
	} else {
		fmt.Println("All is right with the world.")
	}
	//<end id="else"/>

	//<start id="elseif"/>
	grade := 95
	if grade > 90 {
		fmt.Println("Grade : A")
	} else if grade > 80 {
		fmt.Println("Grade : B")
	} else {
		fmt.Println("Study harder, please.")
	}
	//<end id="elseif"/>

	//<start id="simplefor"/>
	total := 1
	for total < 50 {
		total = total + 2
	}
	//<end id="simplefor"/>

	//<start id="for"/>
	for i := 0; //<co id="prestatement" />
		i < 10; //<co id="looptest" />
		i++ {  //<co id="poststatement" />
		fmt.Println(i)
	}
	//<end id="for"/>

	//<start id="nocondition"/>
	for b := 0; ; b++ {
		if b > 10 {
			break
		}
	}
	//<end id="nocondition"/>

	var i int
	i = 10

	//<start id="switchnormal"/>
	switch i {  //<co id="variable" />
	case 5:
		fmt.Println("i is 5")  //<co id="case" />
	case 10:
		fmt.Println("i is 10")
	default:
		fmt.Println("i is neither 5 nor 10")  //<co id="default" />
	}
	//<end id="switchnormal"/>

	//<start id="switchcase"/>
	switch {
	case i < 5:  //<co id="case1" />
		fmt.Println("i is less than 5")
	case i >= 5:  //<co id="case2" />
		fmt.Println("i is 5 or greater")
	}
	//<end id="switchcase"/>

	//<start id="switchdefault"/>
	switch {
	default:
		fmt.Println("i must be 10") //<co id="default" />
	case i > 10:
		fmt.Println("i is greater than 10")
	case i < 10:
		fmt.Println("i is less than 10")

	}
	//<end id="switchdefault"/>

	//<start id="switchexpression"/>
	switch i < 5 {  //<co id="expression" />
	case true:
		fmt.Println("i is less than 5")
	case false:
		fmt.Println("i is greater than 5") 
	}
	//<end id="switchexpression"/>

	showDefer()

	logger.Println("Finished Execution")
}
