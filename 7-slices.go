package main

import "fmt"

/*
	Slices in Go:
		- Dynamic, flexible view into the elements of an array
		- Created using slice literals or the make function
		- Have length and capacity properties
*/
func slices() {
	// Creating a slice using a slice literal, similar to array literal but without size
	intSlice := []int32{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	fmt.Println(len(intSlice)) // length of the slice
	fmt.Println(cap(intSlice)) // capacity of the slice

	intSlice = append(intSlice, 6)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	intSlice2 := []int32{7, 8, 9}
	intSlice = append(intSlice, intSlice2...) // appending another slice, use ... to unpack

	fmt.Println(intSlice)

	length := 3
	capacity := 5

	// Creating a slice using the make function with specified length and capacity
	intSlice3 := make([]int32, length, capacity)
	fmt.Println(intSlice3)
	fmt.Println(len(intSlice3))
	fmt.Println(cap(intSlice3))
}
