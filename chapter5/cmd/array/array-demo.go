package main

import "fmt"

var array [2]int

func main() {

	//var array2 [2][3]int
	array3 := [4][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}

	for i := 0; i < 4; i++ {
		fmt.Println(array3[i])
	}
	fmt.Println(array, array3[1])
}
