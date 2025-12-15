package main

import "fmt"

/*
	For Loops in Go:
		- Used for iterating over collections or executing code repeatedly
		- Three forms: traditional, range-based, and while-like
*/
func forLoop() {
	intArr := [5]int{10, 20, 30, 40, 50}
	for index, value := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	var i = 10

	for i >= 0 {
		if i == 0 {
			break
		}

		fmt.Println(i)
		i--
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
