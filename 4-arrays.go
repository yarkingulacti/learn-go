package main

import "fmt"

/*
	Arrays in Go:
		- Fixed-size collection of elements of the same type
		- Declared with a specific size
		- Elements are accessed via indices
		- Zero-valued by default (e.g., 0 for int, "" for string)
*/
func arrays() {
	var intArr [5]int8 = [5]int8{1, 2, 3, 4, 5}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	intArr[4] = 10
	fmt.Println(intArr[4])

	// array pointers
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}
