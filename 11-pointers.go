package main

import "fmt"

/*
Pointers in Go:
	- Variables that store memory addresses of other variables
	- Declared using the '*' operator
	- The '&' operator is used to get the address of a variable
	- The '*' operator is used to dereference a pointer to access the value at the address
	- Useful for efficient memory usage and modifying variables in functions
*/
func pointers() {
	// var p *int32 // initially nil because it doesn't point to any address
	var p *int32 = new(int32) // allocates memory for an int32 and returns its address
	var i int32 = 31

	fmt.Printf("The value p points to is: %v\n", *p) // dereferencing p gives the zero value of int32, which is 0
	fmt.Printf("The value of i is: %v\n", i)

	*p = 42 // setting the value at the address p points to

	fmt.Printf("After setting, the value p points to is: %v\n", *p)

	// var p2 *int32 // nil pointer
	// *p2 = 21      // This will cause a runtime panic: dereferencing a nil pointer

	fmt.Printf("Before setting, the value of i is: %v\n", i)
	var p3 *int32 = &i // p3 points to the address of i
	*p3 = 21           // setting the value at the address p3 points to, which is i
	fmt.Printf("After setting through pointer, the value of i is: %v\n", i)

	var k int32 = 55
	var j int32

	fmt.Printf("Before Value of k: %v\n", k) // k is 55
	fmt.Printf("Before Value of j: %v\n", j) // j is 0

	j = 99 // modifying j does not affect k

	fmt.Printf("After Value of k: %v\n", k) // k remains 55
	fmt.Printf("After Value of j: %v\n", j) // j is now 99

	var slice = []int32{1, 2, 3}            // slices are reference types
	fmt.Printf("Before Slice: %v\n", slice) // [1 2 3]
	var sliceCopy = slice                   // both slice and sliceCopy point to the same underlying array
	sliceCopy[0] = 99                       // modifying sliceCopy affects slice

	fmt.Printf("After Slice: %v\n", slice) // [99 2 3]

	var array = [3]int32{1, 2, 3}           // arrays are value types
	fmt.Printf("Before Array: %v\n", array) // [1 2 3]
	var arrayCopy = array                   // arrayCopy is a copy of array
	arrayCopy[0] = 99                       // modifying arrayCopy does not affect array

	fmt.Printf("After Array: %v\n", array) // [1 2 3]

	var ages = map[string]int32{
		"Alice": 30,
		"Bob":   25,
	} // maps are reference types
	fmt.Printf("Before Map: %v\n", ages)
	var agesCopy = ages    // both ages and agesCopy point to the same underlying data
	agesCopy["Alice"] = 31 // modifying agesCopy affects ages
	fmt.Printf("After Map: %v\n", ages)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Before square: %v\n", thing1)
	fmt.Printf("The memory location of thing1 array is: %p\n", &thing1)
	var result = square(thing1)
	fmt.Printf("After square: %v\n", result)
	fmt.Printf("The value of thing1 is: %v", thing1)

	var thing2 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nBefore squarePointer: %v\n", thing2)
	fmt.Printf("The memory location of thing2 array is: %p\n", &thing2)
	var result2 = squarePointer(&thing2)
	fmt.Printf("After squarePointer: %v\n", result2)
	fmt.Printf("The value of thing2 is: %v", thing2)
}

// This way we are allocating new memory for the array inside the function
// and returning a new array, leaving the original array unchanged.
// This can cause larger memory usage if the array is large.
// Instead, we could pass a pointer to the array to modify it in place.
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array in square func is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

// This way we are passing a pointer to the array, allowing us to modify
// the original array in place without allocating new memory for a copy.
// This is more memory efficient, especially for large arrays.
func squarePointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array in squarePointer func is: %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
