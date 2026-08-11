package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Question why does it gives me error when i use double quote in reader.ReadString ?

func main() {
	welcome := "Welcome To Pizza Hub"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like to order?")
	order, _ := reader.ReadString('\n')
	fmt.Println("Your order is begin prepared \n orders : ", order)
	fmt.Println("Your order has been serverd! \n would u like to give an rating?")
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your rating: ", rating)
	fmt.Printf("type of rating : %T\n", rating)

	// conversion, converting string to float
	ratingFloat , err  := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("err", err)
	}else {
		fmt.Printf("type of rating : %T\n", ratingFloat)
		
		fmt.Println("ratingFloat", ratingFloat)
	}
}
