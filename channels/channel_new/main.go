package main

import (
	"fmt"
	"net/http"
)


func main(){
	// fmt.Println("test")
	websites := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
	}


	c := make(chan string)

	for _,website := range websites{
		// fmt.Println("webiste name ", website)
		go siteIsUp(website, c)

	}
	for l:= range c{
		go func(link string){
			siteIsUp(link, c)
		}(l)
		
	}
}


func siteIsUp(website string, c chan string){
	// fmt.Println("inside site is up")ls
	_,err := http.Get(website)
	if err != nil{
		fmt.Println("website is down ", website)
		c <- website
	}
	fmt.Println("website is up ", website)
	c <- website
	return
}