package main

import "fmt"

func main (){
	fmt.Println("Pointerrrrrrrrrr")

	// var ptr *int 

	// fmt.Println("ptr", ptr)
	myint := 42
	var ptr = &myint 

	fmt.Println("ptr", ptr)
	fmt.Println("ptr value", *ptr)
	*ptr = *ptr + 10 
	fmt.Println("ptr value after increment", *ptr , myint)
	myint = myint + 10 
	fmt.Println("ptr value after increment", *ptr , myint)
	myint = *ptr + 10 
	fmt.Println("ptr value after increment", *ptr , myint)
	myint = *&myint + 10  //Question : why can i do this,  if we look at right side, it will create a new pointer(value) to myint and then add the & (ram location) to it
	fmt.Println("ptr value after increment", *ptr , myint)

	
	// Question: whats the diffrenc btw *ptr = 52 and  myint = 52
	

}