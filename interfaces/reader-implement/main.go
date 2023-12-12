package main

import (
	"fmt"
	"io"
	"os"
)


func main(){
	fmt.Println(os.Args[1])
	fmt.Println("File read assignment")
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(content)
	fmt.Println("Reading contents of file started")

	io.Copy(os.Stdout, file)
	fmt.Println("File read was successful")
}
