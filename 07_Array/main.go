package main

import "fmt"

func main() {
	fmt.Println("Array")
	var fruit [5]string
	fruit[0] = "apple"
	fruit[1] = "banana"
	fruit[2] = "cherry"
	fruit[3] = "date"
	fruit[4] = "elderberry"
	fmt.Println(fruit)

	vegList := [5]string{"carrot"}
	fmt.Println(vegList)

	
}