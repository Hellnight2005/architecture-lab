package main

import "fmt"

func main() {
	fmt.Println("Maps!")

	// A map stores data in key-value pairs.
	//
	// Syntax:
	// map[KeyType]ValueType
	//
	// Here:
	// Key   -> string
	// Value -> string
	//
	// make() creates and initializes an empty map.
	languages := make(map[string]string)

	// Add key-value pairs to the map.
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:\n", languages)

	// Access a value using its key.
	fmt.Println("JS stands for:", languages["JS"])
	fmt.Println("RB stands for:", languages["RB"])

	// Delete a key-value pair from the map.
	delete(languages, "RB")

	fmt.Println("List of all languages after deleting RB:\n", languages)

	// Loop through a map using range.
	//
	// A map returns two values:
	// key   -> the key
	// value -> the value
	//
	// If we don't need the key, use _ (blank identifier).
	for _, value := range languages {
		fmt.Println("Value is:", value)
	}

	// If we need both key and value:
	//
	// for key, value := range languages {
	//     fmt.Printf("For key %v, the value is %v\n", key, value)
	// }
}
