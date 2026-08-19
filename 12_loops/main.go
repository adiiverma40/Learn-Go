package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops! Loops! Loops! Loops! , ehh, am i looping ")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	fmt.Println("=====================")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("=====================")
	for i := range days {
		fmt.Println(i, days[i]) // it gives index not the value.
	}

	fmt.Println("=====================")
	for index, day := range days {
		fmt.Println(index, day)
	}

	fmt.Println("=====================")

	rogueValue := 1
	for rogueValue < 10 {
		// if rogueValue == 5 {
		// 	break
		// }
		if rogueValue == 5 {
			rogueValue++
			continue // if only continue, it will leave the current loop
			//, starts a new but then value will be same so a while true loop
		}

		if rogueValue == 8 {
			goto lco
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("jumped to line 44")

}
