package main

import "fmt"

func kosullar() {
	if 5 > 3 {
		fmt.Println("5 is greater than 3")
	}

	var sayi int = 10

	if sayi%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	switch sayi {
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
	case sayi < 0:
		fmt.Println("The number is negative")
	case sayi == 0:
		fmt.Println("The number is zero")
	case sayi > 0:
		fmt.Println("The number is positive")
	}
}
