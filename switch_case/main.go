package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){

	fmt.Println("this is program for switch case")
	fmt.Println(os.Args[1:])
	// day := "Saturday"
	day := os.Args[1]
	switch day{
	case "Monday":
		fmt.Println("It's a monday!")
	case "Tuesday":
		fmt.Println("It's a tuesday!")

	case "Saturday":
		fmt.Println("It's a saturday")
	default:
		fmt.Println("Invalid day")
	}
	fmt.Println("Switch program for days ended")

	temperature := os.Args[2]
	temperature2,err := strconv.Atoi(temperature)
	if err != nil{
		fmt.Println("it is an error")
	}
	switch {
	case temperature2 < 5:
		fmt.Println("it is  cold")




	default:
		fmt.Println("it's hot")



	}

}