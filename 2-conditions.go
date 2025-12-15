package main

import "fmt"

func conditions() {
	if 5 > 3 {
		fmt.Println("5 is greater than 3")
	}

	var number int = 10

	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	switch number {
	case 1:
		fmt.Println("The number is one")
	case 5:
		fmt.Println("The number is five")
	case 10:
		fmt.Println("The number is ten")
	default:
		fmt.Println("The number is something else")
	}

	switch {
	case number < 0:
		fmt.Println("The number is negative")
	case number == 0:
		fmt.Println("The number is zero")
	case number > 0:
		fmt.Println("The number is positive")
	}
}
