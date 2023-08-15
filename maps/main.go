package main

import "fmt"

func main(){
	//declaring maps in go
	//first way
	colors := map[string]string{
		"red": "#ff00",
		"green":"#su12",
		"white": "#aa77",		
	}


	//fmt.Println(colors)
	

	//second way
	//var color2 map[string]string
	//third way
	// color3 := make(map[string]string)
	// color3["white"] = "#aa77"
	// delete(color3, "white")
	// fmt.Println(color3)
	printMap(colors)
}

func printMap(c map[string]string){
	for color,hex := range c {
		fmt.Println(color, hex)
	}
}

