package main

import "fmt"

/*
Runes in Go:
	- Alias for int32, represents a Unicode code point
	- Can be indexed to access individual runes
	- Useful for handling characters beyond the ASCII set
*/
func runes() {
	// UTF-8 rune indexing
	var myString = []rune{'r', 'é', 's', 'u', 'm', 'é'}
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed)
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)
}
