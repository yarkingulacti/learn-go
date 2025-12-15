package main

import "fmt"

/*
	Maps in Go:
		- Unordered collection of key-value pairs
		- Keys are unique and used to access values
		- Created using make or map literals
		- Zero-valued by default (e.g., 0 for int, "" for string)
*/
func maps() {
	// Creating maps
	myMap := make(map[string]uint8)
	fmt.Println(myMap)

	// Adding key-value pairs while assigning
	myMap2 := map[string]uint8{"Adam": 30, "Eve": 28}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"]) // zero default value 0 for

	age, exists := myMap2["Sarah"]

	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("Key does not exist")
	}

	// deleting a key-value pair
	delete(myMap2, "Adam")

	fmt.Println(myMap2)

	for key, value := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", key, value)
	}

}
