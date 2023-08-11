package main

import "fmt"

func main(){
	fmt.Println("Hello world")
	intSlice := []int{1,2,3,4,5,6,7,8,9,10}
	for _, ele := range (1,11) {
		if ele%2 == 0{
			fmt.Printf("Number %d is even\n",ele)
		}else {
			fmt.Printf("Number %d is odd\n",ele)
		}
	}
}