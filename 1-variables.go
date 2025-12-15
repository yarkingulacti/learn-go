package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Primitive Data Types in Go:
- Integer Types: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
- Floating-Point Types: float32, float64
- String Type: string
- Rune Type: rune (alias for int32, represents a Unicode code point)
- Boolean Type: bool
*/
func variables() {
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNumUnprecise float32 = 12345678.9 // precision loss
	var floatNumPrecise float64 = 12345678.9   // more precise

	fmt.Println(floatNumUnprecise) // 12345679
	fmt.Println(floatNumPrecise)   // 12345678.9

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + intNum32 // Wrong
	var result float32 = floatNum32 + float32(intNum32) // Correct

	fmt.Println(result)

	var myString string = "Hello, \nGo!"

	fmt.Println(myString)
	fmt.Println(len("ğ"))                    // 2 (2 bytes in UTF-8)
	fmt.Println(utf8.RuneCountInString("ğ")) // 1 (1 rune)

	var myRune rune = 'a' //97 in ASCII
	fmt.Println(myRune)

	var isTrue bool = true
	fmt.Println(isTrue)

	var anotherBool bool // default value is false
	fmt.Println(anotherBool)

	var myString2 string
	fmt.Println(myString2) // default value is empty string

	var myRune2 rune
	fmt.Println(myRune2) // default value is 0

	var myText = "text" // type is inferred as string

	fmt.Println(myText)

	myShort := "short variable declaration" // type inference

	fmt.Println(myShort)

	var var1, var2, var3 int = 1, 2, 3
	fmt.Println(var1, var2, var3)

	const myConst = "This is a constant and cannot be changed. Also, it must be initialized at declaration."

	fmt.Println(myConst)
}
