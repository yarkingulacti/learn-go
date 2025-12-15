package main

import "fmt"

func functions() {
	// common function error handling pattern
	var divisionResult, remainder, err = intDivision(10, 0)

	if err != nil {
		fmt.Printf("Error occurred during division %v", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", divisionResult)
	} else {
		fmt.Printf("The result of the integer division is %v with a remainder of %v", divisionResult, remainder)
	}

	switch {
	case err != nil:
		fmt.Printf("Error occurred during division %v", err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", divisionResult)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v", divisionResult, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The result of the integer division is %v", divisionResult)
	case 1:
		fmt.Printf("The result of the integer division is %v with a remainder of 1", divisionResult)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v", divisionResult, remainder)
	}
}

func intDivision(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, fmt.Errorf("divisor cannot be zero")
	}
	var result int = dividend / divisor
	var remainder int = dividend % divisor
	return result, remainder, nil
}
