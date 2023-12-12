package main

import (
	"errors"
	"fmt"
)

// Divide function returns the result of a/b and an error if b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Creating a new error with the errors.New function.
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	// Example 1: Successful division
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Attempting to divide by zero
	result, err = Divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
