package main

import "fmt"

func main(){

	defer func(){
		if r:=recover(); r!=nil {
			fmt.Println("Panic function recovered", r)	
		}
	}()



	fmt.Println("Start of execution")
	panic("something went wrong!!") 

	fmt.Println("End of execution")




}