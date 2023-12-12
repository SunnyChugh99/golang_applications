package main

import "fmt"


func outer() func(int) int{
	return func(x int) int{
		x++
		return x
	}



}


func adder() func (x int) int{
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}




func main(){
	val := 5
	out := outer()
	add := adder()
	fmt.Println(out(val))
	fmt.Println(add(val))
}

// Function closures is a function value that references variables from outside its body. 
// The function may access and assign values to the referenced variables.

// In Go, a closure is a function value that references variables from outside its body. 
// In other words, a closure allows a function to access variables that are not in its parameter list,
// but are in its lexical scope