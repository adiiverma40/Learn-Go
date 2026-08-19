package main

import (
	"fmt"
	"math/rand"
	"time"
)


// Question:  why does go playground in go.dev always returns 3 in randint? 
func main() {
	fmt.Println("Welcome to dice game!")
	rand.Seed(time.Now().UnixNano()) // if rand.seed is not stored in any variable, and initialized
	// how is it that dice is able to use it
	//
	dice := rand.Intn(6) + 1 // no is always inclusive
	fmt.Println("You rolled a", dice)
	// dice = 8
	switch dice {
	case 1:
		fmt.Println("You can move 1 space")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 spaces")
	case 4:
		fmt.Println("You can move 4 spaces")
	case 5:
		fmt.Println("You can move 5 spaces")
	case 6:
		fmt.Println("You can move 6 spaces and roll the dice again")
	default:
		fmt.Println("Did u used memory cheat? ")
	}

}
