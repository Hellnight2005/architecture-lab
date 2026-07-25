package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	presentTime := time.Now().Nanosecond()
	fmt.Println("presentTime:", presentTime)
	// fmt.Println("Current Time is:", presentTime.Format("02-01-2006 15:04 Monday"))

	createddate := time.Date(2020, time.December, 29, 7, 54, 0, 0, time.UTC)
	fmt.Println("Created Date is:", createddate.Format("02-01-2006 15:04 Monday"))
}
