package main

import "fmt"

// LoginToken is exported because the first letter is uppercase.
// Constants cannot be changed after declaration.
const LoginToken string = "qwertyuiop"

func main() {

	// Explicit type declaration
	var username string = "hellnight2005"

	fmt.Println("Hello from:", username)
	fmt.Printf("Variable type is: %T\n", username)

	// Boolean
	var isLoggedIn bool = true

	fmt.Println("Is user logged in:", isLoggedIn)
	fmt.Printf("Variable type is: %T\n", isLoggedIn)

	// uint8 can store values from 0 to 255
	var smallValue uint8 = 255

	fmt.Println("Small value is:", smallValue)
	fmt.Printf("Variable type is: %T\n", smallValue)

	// Floating-point number
	var smallFloat float64 = 2.356

	fmt.Println("Small float is:", smallFloat)
	fmt.Printf("Variable type is: %T\n", smallFloat)

	// Zero value
	// The default value of int is 0.
	var defaultValue int

	fmt.Println("Default value is:", defaultValue)
	fmt.Printf("Variable type is: %T\n", defaultValue)

	// Type inference
	// Go automatically infers that website is a string.
	var website = "hellnight2005.com"

	fmt.Println("Website is:", website)
	fmt.Printf("Variable type is: %T\n", website)

	// Short variable declaration
	// The type is inferred automatically.
	numberOfUsers := 30000

	fmt.Println("Number of users is:", numberOfUsers)
	fmt.Printf("Variable type is: %T\n", numberOfUsers)

	// Using a constant
	fmt.Println("Login token is:", LoginToken)
	fmt.Printf("Variable type is: %T\n", LoginToken)
}
