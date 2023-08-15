package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipcode int
}

//defining structs
type person struct {
	firstName string
	lastName string
	contactInfo 
}

func main(){
	//First way of declaring structs
	// alex := person{"Alex","Anderson"}
	// fmt.Println(alex)

	// //second way
	// Jason := person{firstName: "Jason", lastName: "Anderson"}
	// fmt.Println(Jason)	

	// //third way

	// var john person
	// john.firstName = "John"
	// john.lastName = "Cena"

	jim := person {
		firstName: "jim",
		lastName: "party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipcode: 421004,
		},
	}
	//value can be converted to address by ==>  &value
	jimPointer := &jim
	jimPointer.updateFirstName("James")
	jim.print()

	sliceString := []string{"Hi","How", "Are", "You"}
	replaceFirst(sliceString)
	fmt.Println(sliceString)

}
// address can be converted to value by  ==> *address
func (pointerToPerson *person) updateFirstName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}

func replaceFirst(sliceString []string){
	sliceString[0] = "Bye"

}