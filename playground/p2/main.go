package main

import (
	"fmt"
	"reflect"
)

func printTypeAndValue(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("Type: %v\n", t)
	fmt.Printf("Value: %v\n", v)
}

func main() {
	var i int = 42
	var f float64 = 3.14
	var s string = "Hello, Go!"

	printTypeAndValue(i)
	printTypeAndValue(f)
	printTypeAndValue(s)
}
