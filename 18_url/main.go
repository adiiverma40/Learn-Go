package main

import (
	"fmt"
	"net/url"
)

// why does compilers have problem with unused import of variables, i didnt get any errors in intrepretter
	// "golang.org/x/text/unicode/rangetable"


const myurl string = "http://localhost:8000/update/release?coursename=reactjs&paymentid=ghasasdf"


func main(){
	fmt.Println("Handling urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Hostname())
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)

	qprams := result.Query()
	fmt.Printf("type of query %T\n", qprams)
	fmt.Println(qprams["coursename"])

	for _,value := range qprams {
		fmt.Println("param is : ", value)
	}
	fmt.Println("==============Construct url=============")
// why to pass refrence instead of copy? 
	parstOfurl := &url.URL{
		Scheme: "http",
		Host: "localhost:8000", //why is there no port options?
		Path: "release/updated",
		
		
	}

	fmt.Println(parstOfurl.String())
}

