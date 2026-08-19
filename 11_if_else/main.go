package main

import "fmt"

func main() {
	fmt.Println("If Else")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {

		result = "Watch out"
	} else {
		result = "exactly 10 "
	}

	fmt.Println(result)
	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	// Asigning and checking in the go 
	// example, web request 
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than or equal to 10")
	}
}