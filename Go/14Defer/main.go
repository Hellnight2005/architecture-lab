package main

import (
	"fmt"
)

// multiple stateme of the defer it use the lIFO last in first out to execute the defer statement
func main() {
	fmt.Println("Defer!")
	defer fmt.Println("This is the third deferred statement.")
	defer fmt.Println("This is the first deferred statement.")
	fmt.Println("This is the second statement.")
	myDferFunc()
}

// now all thing normale as it rerach the ender it for the run the function then the defer statement will be executed in the reverse order of their appearance in the code.

func myDferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
