package main

import (
	"fmt"
	"os/exec"
)

/*

//<start id="declaration"/>
func printString(printMe string  <co id="inputType" />) {
	fmt.Println(printMe)
}

func plusOne(num int <co id="inputType2" />) int <co id="returnType" />{
	return num + 1
}

//<end id="declaration"/>

*/

//<start id="retval"/>
func stringLength(in string) int {
	return len(in)
}

//<end id="retval"/>

//<start id="namedreturn"/>
func stringLengthAgain(in string) (count int) { //<co id="nameddeclare" />
	count = len(in) //<co id="usingnd" />
	return
}

//<end id="namedreturn"/>

//<start id="multireturn"/>
func ParseBool(str string) (value bool, err error) { //<co id="multireturns" />
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil //<co id="multireturnsTrue" />
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil //<co id="multireturnsFalse" />
	}
	return false, syntaxError("ParseBool", str) //<co id="multireturnsError" />
}

//<end id="multireturn"/>

//<start id="errorreturn"/>
func runOSCommand(command string) (out []byte, err error) {
	out, err = exec.Command(command).Output()
	return
}

//<end id="errorreturn"/>

//<start id="variadic"/>
func printStringValues(values ...string) {
	for i, v := range values {
		fmt.Println("Value:", v, "Ordinal:", i)
	}
}

//<end id="variadic"/>
func getExternalStatus() string {
	return "GOOD"
}

//<start id="nilnamed"/>
func checkStatus() (status string, err error) {
	status = getExternalStatus()
	if status == "ERROR" {
		err = errors.New("External Status returned error")
	}
	return
}

//<end id="nilnamed" />

//<start id="funcparm"/>
func checkStatus(
	status string,
	errorFunc func(string), //<co id="funcasparm1" />
	successFunc func()) { //<co id="funcasparm2" />
	if status == "FAIL" {
		errorFunc(status)
		return
	}
	successFunc()
}

//<end id="funcparm"/>

func main() {

	// START:callvariadic
	string1 := "first"
	string2 := "second"
	string3 := "third"

	printStringValues(string1)
	printStringValues(string1, string2)
	printStringValues(string1, string2, string3)
	// END:callvariadic

	//<start id="functiontype"/>
	var errorFunction = func(s string) { //<co id="errorfunction" />
		fmt.Printf("Error occurred: %s\n", s)
	}
	var successFunction = func() { //<co id="successfunction" />
		fmt.Printf("Success.")
	}

	checkStatus("FAIL", errorFunction, successFunction)    //<co id="checkfailed" />
	checkStatus("SUCCESS", errorFunction, successFunction) //<co id="checksuccess" />
	//<end id="functiontype"/>

	//<start id="anonymous"/>
	checkStatus("FAIL",
		func(s string) { fmt.Printf("Error occurred: %s\n", s) },
		func() { fmt.Printf("Success.") },
	)
	//<end id="anonymous"/>

	//<start id="closure"/>

	string4 := "Gophers Were Here" //<co id="outsidevar" />

	go func() {
		fmt.Println("Enclosed " + string4) //<co id="enclosing" />
	}() //<co id="parens" />
	//<end id="closure"/>

	//<start id="checkerror"/>
	b, e := strconv.ParseBool("I'm not a bool")
	if e != nil {
		//there was an error!
		fmt.Println(e)
	}

	// This will print false
	fmt.Println(b)
	//<end id="checkerror"/>

}
