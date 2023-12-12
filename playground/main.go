// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)
func main() {
	fmt.Println("Hello World")
	mySlice := make([]int, 3, 5)
	fmt.Println(mySlice)


	m := make(map[string]int)

	c := make(chan string)
	fmt.Println(c)
	fmt.Println(m)

	myString := "42"

	myInt,err := strconv.Atoi(myString)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("The variable is: ",myInt, "having data type ", reflect.TypeOf(myInt))


	myString2 := strconv.Itoa(myInt)
	fmt.Println("The variable is ", myString2, " having data type ", reflect.TypeOf(myString2))




}
