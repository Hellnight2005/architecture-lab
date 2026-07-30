package main

import "fmt"

func main() {

	// Declare a pointer
	// The default value of ptr is nil.
	var ptr *int
	fmt.Println("Value of pointer is:", ptr)

	// Create a variable
	mynumber := 23

	// & gets the memory address of mynumber
	// ptr stores the address of mynumber
	ptr = &mynumber

	// Print the memory address stored in ptr
	fmt.Println("Value of pointer is:", ptr)

	// *ptr dereferences the pointer
	// It gives the value stored at the address
	fmt.Println("Value of the variable is:", *ptr)

	// Modify the original variable through the pointer
	*ptr = *ptr + 2

	// mynumber is now 25
	fmt.Println("Value of the variable is:", mynumber)
}
