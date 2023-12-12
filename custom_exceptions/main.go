package main

import (
	"fmt"
)

// CustomError is a custom error type that includes additional information.
type CustomError struct {
	Code    int
	Message string
}

// Implementing the error interface for CustomError.
func (e CustomError) Error() string {
	return fmt.Sprintf("Error11 %d: %s", e.Code, e.Message)
}

// Divide function returns the result of a/b and a custom error if b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Creating a custom error with additional information.
		return 0, CustomError{Code: 400, Message: "division by zero is not allowed"}
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

	// Example 3: Using information from the custom error
	if customErr, ok := err.(CustomError); ok {
		fmt.Printf("CustomError - Code: %d, Message: %s\n", customErr.Code, customErr.Message)
	}
}
