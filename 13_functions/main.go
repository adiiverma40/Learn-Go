package main

import "fmt"

// Why is main functions executes first without being called
func main() {
	Greater()
	fmt.Println("Welcome to functions!")
	GreaterTwo()
	result := ReallyHardAddition(5, 10)
	fmt.Println("Result of really hard addition that cant be done manully is : ", result)
	reallyReallyHardAddition := IcantCallReallyHardAdditionEverytime(2,3,4,5,5,6)
	fmt.Println("Result of really hard addition that cant be done manully and cant be done by calling adding function again and again :", reallyReallyHardAddition)

}


func IcantCallReallyHardAdditionEverytime(values ...int) int{
	sum := 0
	for _, v := range values{
		sum += v
	}
	return sum
}


func ReallyHardAddition(a int, b int) int {
	return a + b
}


func GreaterTwo(){
	fmt.Println("Hello x Again!")
}

// Anonymous func
// func (){
// 	fmt.Println("Hello x!")
// }()

func Greater(){
	fmt.Println("Hello x!")
}
