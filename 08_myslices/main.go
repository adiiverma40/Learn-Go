package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SLICESSSSSSSSS!")

	var fruits = []string{"1", "2", "3"}
	fmt.Printf("fruits %T\n", fruits)
	fmt.Println(fruits)
	fruits = append(fruits, "4", "5")
	// fruits= append(fruits, "mango", "banana", 1 , 2)
	// why cant i append int into slices
	//
	fmt.Println(fruits)
	// anotherSlice := []string{fruits[0]}  // result 1
	anotherSlice := []string{}
	// anotherSlice = append(anotherSlice, fruits)
	anotherSlice = append(anotherSlice, fruits...) // spread like butter

	fmt.Println("AnotherSlice", anotherSlice)
	fruits = append(fruits[1:])
	fmt.Println(fruits)
	// fmt.Println(fruits)
	fruits = append(fruits[1:3]) //[mango tomato mango banana]   range are no inlcusive.
	fmt.Println(fruits)
	fruits = append(fruits[:3]) // If range is not inclusive then why does it inlcued the 3rd element?
	// Answer: it did not, here counting starts from 0 so, 3 = 4. so 4th element is not included in output
	// //[3 4 5] , why does it have 3, 4, 5. not 1, 2 ,3? did u forget the previous operations
	fmt.Println(fruits)
	fmt.Println("==============================")

	score := make([]int, 4)
	score[0] = 1
	score[1] = 2
	score[2] = 3
	score[3] = 4
	// score[4] = 23245342
	score = append(score, 5, 6, 7, 8) // why define as array, and then make it slices?/
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	fmt.Println("==============================")
	fmt.Println("Remove value from slices based on index ")
	fmt.Println("==============================")

	var index int = 3

	score = append(score[:index])

}

// func main() {
// 	fmt.Println("SLICESSSSSSSSS!")

// 	var fruits = []string{"apple", "mango", "tomato"}
// 	fmt.Printf("fruits %T\n", fruits)
// 	fmt.Println(fruits)
// 	fruits= append(fruits, "mango", "banana")
// 	// fruits= append(fruits, "mango", "banana", 1 , 2)
// 	// why cant i append int into slices
// 	//
// 	fmt.Println(fruits)
// 	fruits = append(fruits[1:])
// 	fmt.Println(fruits)
// 	fmt.Println(fruits)
// 	fruits = append(fruits[1:3]) //[mango tomato mango banana]
// 	fmt.Println(fruits)

// }
