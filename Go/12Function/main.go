package main

import "fmt"

func main() {
	fmt.Println("Functions!")

	fmt.Println("Addition:", add(5, 3))
	fmt.Println("Subtraction:", subtract(5, 3))

	// Create a slice with capacity for 101 integers.
	values := make([]int, 0, 101)

	// Add numbers from 0 to 100.
	for i := 0; i <= 100; i++ {
		values = append(values, i)
	}

	// Pass all values to the variadic function.
	result := proAdd(values...)

	fmt.Println("Total:", result)
}

// add returns the sum of x and y.
func add(x int, y int) int {
	return x + y
}

// subtract returns the difference between x and y.
func subtract(x int, y int) int {
	return x - y
}

// proAdd accepts any number of integers and returns their total.
func proAdd(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}
