package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome! lets practice time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(presentTime.Format("01-02-2006 monday"))

	fmt.Println(presentTime.Format("2006-01-02 monday"))
	fmt.Println(presentTime.Format("01/02/2006 Mon"))

	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Mon"))
	fmt.Println(presentTime.Format("01/02/2006 05:04:15 Mon"))

	createDate := time.Date(2026, time.December , 32, 20 , 20 ,20 , 0, time.UTC)
	fmt.Println(createDate.Format("01/02/2006 15:04:05 Mon"))
	// Question: why didnt it give me error when i gave 32 as the day?
	// Instead of error why did the date moved up by 1? 32 - 31 = 1?


	
}
