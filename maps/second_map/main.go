package main

import (
	"fmt"
	"strconv"
)

func main(){

	var init map[string]int

	init2 := make(map[string]int)

	colors := map[string]int{
		"red": 1234,
		"yellow": 899,
	}

	fmt.Println(colors)
	fmt.Println(init)
	fmt.Println(init2)

	colors["blue"] = 4545
	colors["green"] = 1

	delete(colors, "blue")
	fmt.Println(colors)
	fmt.Println("proper seq of map looping")

	for color, hex := range colors{
		fmt.Println("Color is " + color + " hex is " + strconv.Itoa(hex))
	}


}