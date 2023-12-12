package main

import "fmt"

func main(){
	fmt.Println("Map ds detailing in go")
	var dataMap map[string]interface{}
	// make(map[string]int)
	
	dataMap = make(map[string]interface{})
	fmt.Println(dataMap)
}