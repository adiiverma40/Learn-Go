package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:8000/update/release"

func main(){
	fmt.Println("Web Request")

	res , err := http.Get(url)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("Response type is : %T\n", res)
	defer res.Body.Close() // always close the connection

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(dataBytes))
}