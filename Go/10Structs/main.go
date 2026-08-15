package main

import "fmt"

// ============================================================
// STRUCTS IN GO
// ============================================================
//
// A struct (structure) is a user-defined data type that groups
// related data of different types into a single value.
//
// Go does NOT use classes like Java/C++.
// Instead, Go commonly uses:
//   - structs for grouping data
//   - functions/methods for behavior
//
// Think of a struct as a template/blueprint for creating
// values that contain multiple related fields.
//
// Example:
// A User can have:
//   Name   -> string
//   Email  -> string
//   Status -> bool
//   Age    -> int
//
// ============================================================

func main() {
	fmt.Println("Structs!")

	// Create a User struct value.
	//
	// The values are assigned according to the order
	// of the fields defined in the User struct:
	//
	// Name   -> "Abhijeet"
	// Email  -> "abhijeet@gmail.com"
	// Status -> true
	// Age    -> 25
	//
	// This is called an unkeyed struct literal.
	abhijeet := User{"Abhijeet", "abhijeet@gmail.com", true, 25}

	fmt.Println("User details are:", abhijeet)

	// Access individual struct fields using the dot (.) operator.
	fmt.Println("User name is:", abhijeet.Name)
	fmt.Println("User email is:", abhijeet.Email)
	fmt.Println("User status is:", abhijeet.Status)
	fmt.Println("User age is:", abhijeet.Age)

	// --------------------------------------------------------
	// KEYED STRUCT LITERAL
	// --------------------------------------------------------
	//
	// A better and more readable way to create a struct is
	// to explicitly specify the field names.
	//
	// This avoids depending on the order of fields.

	user := User{
		Name:   "Rahul",
		Email:  "rahul@gmail.com",
		Status: true,
		Age:    22,
	}

	fmt.Println("User details are:", user)

	// --------------------------------------------------------
	// MODIFY STRUCT FIELDS
	// --------------------------------------------------------
	//
	// Struct fields can be changed using the dot operator.

	user.Name = "Rahul Sharma"
	user.Age = 23

	fmt.Println("Updated user:", user)

	// --------------------------------------------------------
	// STRUCT WITH ZERO VALUES
	// --------------------------------------------------------
	//
	// If we create a struct without providing values,
	// Go automatically assigns the zero value for each field.
	//
	// string -> ""
	// bool   -> false
	// int    -> 0

	var emptyUser User

	fmt.Println("Empty user:", emptyUser)
	fmt.Println("Empty user name:", emptyUser.Name)
	fmt.Println("Empty user age:", emptyUser.Age)

	// --------------------------------------------------------
	// STRUCT POINTER
	// --------------------------------------------------------
	//
	// A pointer can store the memory address of a struct.
	//
	// &user -> gets the address of user
	// *userPtr -> accesses the value stored at that address

	userPtr := &user

	fmt.Println("User through pointer:", userPtr)

	// Go automatically dereferences struct pointers when
	// accessing fields.
	//
	// These two are equivalent:
	//
	// (*userPtr).Name
	// userPtr.Name

	fmt.Println("Name through pointer:", userPtr.Name)

	// We can also modify the original struct through a pointer.
	userPtr.Age = 24

	fmt.Println("Updated age:", user.Age)
}

// ============================================================
// DEFINING A STRUCT
// ============================================================
//
// Syntax:
//
// type StructName struct {
//     FieldName FieldType
//     FieldName FieldType
// }
//
// The User struct below defines the fields that every User
// value can contain.
//
// ============================================================

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// ============================================================
// EXPORTED VS UNEXPORTED FIELDS
// ============================================================
//
// Go uses capitalization to control visibility between packages.
//
// Field starting with CAPITAL letter:
//     Name
//
// -> Exported
// -> Can be accessed from another package.
//
// Field starting with SMALL letter:
//     name
//
// -> Unexported
// -> Cannot be accessed directly from another package.
//
// Example:
//
// type User struct {
//     Name  string // Exported
//     email string // Unexported
// }
//
// Another package can do:
//
// user.Name
//
// But it cannot do:
//
// user.email
//
// The same rule also applies to types, functions, variables,
// constants, etc.
//
// Capital letter = exported
// Small letter   = unexported
//
// ============================================================
//
// IMPORTANT:
// Structs are NOT classes.
//
// Go does not have traditional classes and inheritance like
// Java or C++.
//
// Instead, Go uses structs to represent data and methods to
// attach behavior to those structs.
//
// Example:
//
// func (u User) GetName() string {
//     return u.Name
// }
//
// This is called a METHOD.
//
// ============================================================

// A struct can also contain fields of many different types.
//
// Example:
//
// type Product struct {
//     Name     string
//     Price    float64
//     Quantity int
//     InStock  bool
// }
//
// Here one struct combines four different data types.
//
// ============================================================
//
// QUICK SUMMARY
//
// Struct:
//     Groups related data together.
//
// Define:
//     type User struct { ... }
//
// Create:
//     user := User{"Abhijeet", "email@gmail.com", true, 25}
//
// Recommended creation:
//     user := User{
//         Name:   "Abhijeet",
//         Email:  "email@gmail.com",
//         Status: true,
//         Age:    25,
//     }
//
// Access field:
//     user.Name
//
// Modify field:
//     user.Name = "Rahul"
//
// Pointer:
//     userPtr := &user
//
// Access through pointer:
//     userPtr.Name
//
// Exported field:
//     Name
//
// Unexported field:
//     name
//
// Zero values:
//     string -> ""
//     int    -> 0
//     bool   -> false
//
// ============================================================
