package main

import (
	"fmt"
	"strings"
)

/*
Strings in Go:
  - Immutable sequence of bytes
  - UTF-8 encoded by default
  - Can be indexed to access individual bytes
  - Length can be obtained using len() function
  - Concatenation can be done using + operator or strings.Builder for efficiency
*/
func funcStrings() {
	// UTF-8 byte indexing
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed)
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	fmt.Println("\nThe length of 'myString' is %v", len(myString))

	var strSlice []string = []string{"r", "é", "s", "u", "m", "é"}
	var catStr string

	for i := range strSlice {
		// create a new string each time
		catStr += strSlice[i]
	}

	// catStr[0] = 'R' // strings are immutable

	fmt.Println(catStr)

	var stringBuilder strings.Builder
	var strSlice2 []string = []string{"r", "é", "s", "u", "m", "é"}

	for i := range strSlice2 {
		stringBuilder.WriteString(strSlice2[i])
	}

	fmt.Println(stringBuilder.String())
}
