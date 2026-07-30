package main

import (
	"fmt"
	"sort"
)

func main() {

	// ======================================================
	// 1. Declaring an Empty Slice
	// ======================================================

	var fruitList []string

	fmt.Printf("Type of fruitList: %T\n", fruitList)
	fmt.Println("Initial fruit list:", fruitList)
	fmt.Println("Length:", len(fruitList))
	fmt.Println("Capacity:", cap(fruitList))

	// ======================================================
	// 2. Adding Elements using append()
	// ======================================================

	fruitList = append(fruitList, "Apple")
	fruitList = append(fruitList, "Banana")
	fruitList = append(fruitList, "Grapes")

	// Multiple values can be appended at once
	fruitList = append(fruitList, "Mango", "Orange", "Pineapple")

	fmt.Println("\nFruit list after append:", fruitList)
	fmt.Println("Length:", len(fruitList))
	fmt.Println("Capacity:", cap(fruitList))

	// ======================================================
	// 3. Creating a Sub-slice
	// ======================================================

	subFruitList := fruitList[1:4]

	fmt.Println("\nOriginal Fruit List:", fruitList)
	fmt.Println("Sub Slice:", subFruitList)

	// ======================================================
	// 4. Creating a Slice using make()
	// ======================================================

	highScores := make([]int, 4)

	highScores[0] = 1052
	highScores[1] = 500
	highScores[2] = 330
	highScores[3] = 420

	fmt.Println("\nHigh Scores:", highScores)

	// ======================================================
	// 5. Calculating the Total
	// ======================================================

	var total int

	for i := 0; i < len(highScores); i++ {
		total += highScores[i]
	}

	fmt.Println("Total High Score:", total)

	// ======================================================
	// 6. Sorting a Slice
	// ======================================================

	sort.Ints(highScores)

	fmt.Println("Sorted High Scores:", highScores)

	// ======================================================
	// 7. Removing an Element by Index
	// ======================================================

	colors := []string{"Red", "Green", "Blue", "Yellow", "Purple"}

	fmt.Println("\nOriginal Colors:", colors)

	index := 2 // Remove "Blue"

	colors = append(colors[:index], colors[index+1:]...)

	fmt.Println("After Removing Index 2:", colors)
}
