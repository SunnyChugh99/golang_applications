package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main(){

	fmt.Println("HTTP")
	resp,err := http.Get("http://google.com")
	if err != nil{
		fmt.Println("ERROR: ",err)
		os.Exit(1)
	}

	//fmt.Println(resp)
	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	//fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)


}	

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("You just wrote these bytes-",len(bs))
	return len(bs), nil
}