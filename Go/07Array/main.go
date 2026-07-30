package main

import "fmt"

func main() {

	// Declaring an array of 4 strings
	var fruitList [4]string

	// Assigning values using indexes
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[2] = "Grapes" // Not assigned, so it remains the zero value ("")
	fruitList[3] = "Mango"

	// Printing the array
	fmt.Println("Fruit list is:", fruitList)

	// Printing the length of the array
	fmt.Println("Fruit list length is:", len(fruitList))

	// Declaring and initializing an array using an array literal
	var veg = [3]string{"Carrot", "Potato", "Tomato"}

	// Printing the initialized array
	fmt.Println("Vegetable list is:", veg)

	// Printing the length of the initialized array
	fmt.Println("Vegetable list length is:", len(veg))
}
