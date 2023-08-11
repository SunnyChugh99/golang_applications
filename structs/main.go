package main

import (
	"fmt"
)

//defining structs
type person struct {
	firstName string
	lastName string
}

func main(){
	//First way of declaring structs
	alex := person{"Alex","Anderson"}
	fmt.Println(alex)

	//second way
	Jason := person{firstName: "Jason", lastName: "Anderson"}
	fmt.Println(Jason)	

}