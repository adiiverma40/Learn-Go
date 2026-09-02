package main

import "fmt"

func main(){
	defer fmt.Println("world") // change the line from 6 to just before the final closing bracket,
	//but why do we need it?
	// LIFO: Last in first out
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello, Lets learn differ")
	myDefer()

}



func myDefer(){
	for i := 0; i <5; i++{
		defer fmt.Println(i)
	}
	//main() // why does this creates a loop of 
	// fmt.Println("Hello, Lets learn differ") and not give circular import error 
}