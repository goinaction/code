package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	slice := array[1:6]

	slice1 := make([]int, 2)
	slice1 = slice
	slice = append(slice1, 22, 55, 66, 77, 88)
	fmt.Println(slice)
}
