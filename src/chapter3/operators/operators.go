package main

import (
	"fmt"
)

func main() {
	//<start id="equal"/>
	fmt.Println("1 == 2", 1 == 2)
	//<end id="equal"/>

	//<start id="equalstring"/>
	one := "1"
	fmt.Println("'1' == '2'", one == "2")
	two := "2"
	fmt.Println("'2' == '2'", two == "2")
	//<end id="equalstring"/>

	//<start id="equalstruct"/>
	type User struct {
		FirstName    string
		LastName     string
		EmailAddress string
	}
	g1 := User{"Gary", "Gopher", "gary@golang.org"}
	g2 := User{"Gary", "Gopher", "gary@golang.org"}
	g3 := User{"Gene", "Gopher", "gary@golang.org"}
	//Compare the equal structs
	fmt.Println("g1==g2", g1 == g2) //<co id="comparestructequal" />
	//Compare the inequal structs
	fmt.Println("g2==g3", g2 == g3) //<co id="comparestructinequal" />
	//<end id="equalstruct"/>

	//<start id="notequal"/>
	fmt.Println("1 == 2", 1 != 2)
	//<end id="notequal"/>

	//<start id="lessthan"/>
	fmt.Println("1 < 2", 1 < 2)
	//<end id="lessthan"/>

	//<start id="logicaland"/>
	signedIn := true
	verifiedEmail := false
	if signedIn && verifiedEmail {
		fmt.Println("Verification complete.")
	} else {
		fmt.Println("Verification incomplete.")
	}
	//<end id="logicaland"/>

	//<start id="logicalor"/>
	signedIn = true
	verifiedCookie := false
	if signedIn || verifiedCookie {
		fmt.Println("OK to allow access to settings")
	}
	//<end id="logicalor"/>

	//<start id="logicalnot"/>
	verifiedPhone := false
	if !verifiedPhone  {
		fmt.Println("Display phone verification reminder.")
	}
	//<end id="logicalnot"/>

}
