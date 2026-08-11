package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question why does it gives me error when i use double quote in reader.ReadString ? 

func main() {
	welcome := "Welcome To Pizza Hub"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like to order?")
	order, _ := reader.ReadString('\n')
	fmt.Println("Your order is begin prepared \n orders : ", order)
}
