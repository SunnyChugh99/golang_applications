package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	fmt.Println("channels task")

	links := []string{
		"http://google.com",
		"http:/facebook.com",
		"http:/stackoverflow.com",
		"http:/golang.org",
		"http:/amazon.com",
	}

	c := make(chan string)

	for _,link := range links{
		go checkLink(link, c)

	}	

	//wait for the value to be sent over to channel. When we get, log it out immediately
	// for i :=0; i < len(links); i++{	
	// fmt.Println(<-c)
	// // }

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l:= range c {
		//time.Sleep(5 * time.Second)
		//go checkLink(l, c)
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}


func checkLink(link string, c chan string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		
		// send the value 'Might be down i think' to the channel
		c <- link
		return
	}
	fmt.Println(link, " is up!")	
	c <- link
}
