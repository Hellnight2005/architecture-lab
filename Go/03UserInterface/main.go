package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Hell0 user"
	fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")

	// common ok || err ok
	// issuse here that all the thing read by the Reader is awalyes the string
	// name, _ := reader.ReadString('\n')
	// fmt.Println("Hello", name)
	// fmt.Printf("Type of name is: %T\n", name)

	fmt.Println("can you rate the pizza from 1  to 5")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')
	// here we conversion the string to int
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to the rating", numRating+1)

	}

	// fmt.Println("You rated the pizza", numRating)
	// fmt.Printf("Type of numRating is: %T\n", numRating)
	// fmt.Printf("Type of rating is: %T\n", rating)
}
