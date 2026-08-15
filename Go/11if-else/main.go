package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================================
// CONDITIONS IN GO
// ============================================================
//
// Go provides several ways to make decisions:
//
// 1. if
// 2. if / else
// 3. if / else if / else
// 4. switch
//
// Go does NOT require parentheses around conditions.
//
// Correct:
//     if age >= 18 {
//
// Not:
//     if (age >= 18) {
//
// ============================================================

func main() {
	fmt.Println("Conditions in Go")

	// bufio.Reader is used to read input from the terminal.
	reader := bufio.NewReader(os.Stdin)

	// ========================================================
	// PROGRAM 1: USER LOGIN LEVEL
	// ========================================================
	//
	// Take the number of logins from the user and classify
	// the user based on the number of logins.
	//
	// < 10  -> Novice
	// < 20  -> Intermediate
	// < 30  -> Advanced
	// >= 30 -> Expert
	//
	// ========================================================

	fmt.Print("Enter login count: ")

	loginCountStr, _ := reader.ReadString('\n')

	// ReadString returns a string.
	// TrimSpace removes spaces and the newline character.
	loginCountStr = strings.TrimSpace(loginCountStr)

	// Atoi converts a string into an integer.
	loginCount, _ := strconv.Atoi(loginCountStr)

	var result string

	if loginCount < 10 {
		result = "Novice User"
	} else if loginCount < 20 {
		result = "Intermediate User"
	} else if loginCount < 30 {
		result = "Advanced User"
	} else {
		result = "Expert User"
	}

	fmt.Println("Result:", result)

	// ========================================================
	// PROGRAM 2: GRADE CALCULATOR
	// ========================================================
	//
	// Grade based on marks:
	//
	// 90 - 100 -> A
	// 80 - 89  -> B
	// 70 - 79  -> C
	// 60 - 69  -> D
	// Below 60 -> F
	//
	// ========================================================

	fmt.Print("\nEnter your marks: ")

	marksStr, _ := reader.ReadString('\n')
	marksStr = strings.TrimSpace(marksStr)

	marks, _ := strconv.Atoi(marksStr)

	var grade string

	if marks >= 90 && marks <= 100 {
		grade = "A"
	} else if marks >= 80 {
		grade = "B"
	} else if marks >= 70 {
		grade = "C"
	} else if marks >= 60 {
		grade = "D"
	} else if marks >= 0 {
		grade = "F"
	} else {
		grade = "Invalid marks"
	}

	fmt.Println("Grade:", grade)

	// ========================================================
	// PROGRAM 3: EVEN, ODD, OR ZERO
	// ========================================================
	//
	// The modulus operator (%) returns the remainder.
	//
	// number % 2 == 0 -> Even
	// number % 2 != 0 -> Odd
	//
	// Zero is technically even, but we handle it separately
	// here for demonstration.
	//
	// ========================================================

	fmt.Print("\nEnter a number: ")

	numberStr, _ := reader.ReadString('\n')
	numberStr = strings.TrimSpace(numberStr)

	number, _ := strconv.Atoi(numberStr)

	if number == 0 {
		fmt.Println("The number is Zero")
	} else if number%2 == 0 {
		fmt.Println("The number is Even")
	} else {
		fmt.Println("The number is Odd")
	}

	// ========================================================
	// PROGRAM 4: POSITIVE, NEGATIVE, OR ZERO
	// ========================================================

	fmt.Print("\nEnter another number: ")

	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)

	num, _ := strconv.Atoi(numStr)

	if num > 0 {
		fmt.Println("Positive number")
	} else if num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}

	// ========================================================
	// PROGRAM 5: AGE CHECK
	// ========================================================
	//
	// Demonstrates logical operators.
	//
	// && -> AND
	// || -> OR
	// !  -> NOT
	//
	// ========================================================

	fmt.Print("\nEnter your age: ")

	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)

	age, _ := strconv.Atoi(ageStr)

	if age < 0 {
		fmt.Println("Invalid age")
	} else if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teenager")
	} else if age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior Citizen")
	}

	// ========================================================
	// PROGRAM 6: VOTING ELIGIBILITY
	// ========================================================

	fmt.Print("\nEnter age for voting check: ")

	voteAgeStr, _ := reader.ReadString('\n')
	voteAgeStr = strings.TrimSpace(voteAgeStr)

	voteAge, _ := strconv.Atoi(voteAgeStr)

	if voteAge >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// ========================================================
	// PROGRAM 7: LARGEST OF TWO NUMBERS
	// ========================================================

	fmt.Print("\nEnter first number: ")

	firstStr, _ := reader.ReadString('\n')
	firstStr = strings.TrimSpace(firstStr)

	first, _ := strconv.Atoi(firstStr)

	fmt.Print("Enter second number: ")

	secondStr, _ := reader.ReadString('\n')
	secondStr = strings.TrimSpace(secondStr)

	second, _ := strconv.Atoi(secondStr)

	if first > second {
		fmt.Println("First number is larger")
	} else if second > first {
		fmt.Println("Second number is larger")
	} else {
		fmt.Println("Both numbers are equal")
	}

	// ========================================================
	// PROGRAM 8: LARGEST OF THREE NUMBERS
	// ========================================================

	fmt.Print("\nEnter third number: ")

	thirdStr, _ := reader.ReadString('\n')
	thirdStr = strings.TrimSpace(thirdStr)

	third, _ := strconv.Atoi(thirdStr)

	if first >= second && first >= third {
		fmt.Println("First number is the largest")
	} else if second >= first && second >= third {
		fmt.Println("Second number is the largest")
	} else {
		fmt.Println("Third number is the largest")
	}

	// ========================================================
	// PROGRAM 9: CHECK DIVISIBILITY
	// ========================================================
	//
	// Check whether a number is divisible by:
	// 3
	// 5
	// both 3 and 5
	//
	// ========================================================

	fmt.Print("\nEnter a number for divisibility check: ")

	divStr, _ := reader.ReadString('\n')
	divStr = strings.TrimSpace(divStr)

	divNumber, _ := strconv.Atoi(divStr)

	if divNumber%3 == 0 && divNumber%5 == 0 {
		fmt.Println("Divisible by both 3 and 5")
	} else if divNumber%3 == 0 {
		fmt.Println("Divisible by 3")
	} else if divNumber%5 == 0 {
		fmt.Println("Divisible by 5")
	} else {
		fmt.Println("Not divisible by 3 or 5")
	}

	// ========================================================
	// PROGRAM 10: SWITCH STATEMENT
	// ========================================================
	//
	// switch is useful when comparing one value against
	// multiple possible values.
	//
	// Syntax:
	//
	// switch value {
	// case value1:
	//     ...
	// case value2:
	//     ...
	// default:
	//     ...
	// }
	//
	// ========================================================

	fmt.Print("\nEnter a number from 1 to 7: ")

	dayStr, _ := reader.ReadString('\n')
	dayStr = strings.TrimSpace(dayStr)

	day, _ := strconv.Atoi(dayStr)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// ========================================================
	// PROGRAM 11: CALCULATOR USING SWITCH
	// ========================================================
	//
	// Demonstrates switch with operators.
	//
	// + -> Addition
	// - -> Subtraction
	// * -> Multiplication
	// / -> Division
	//
	// ========================================================

	fmt.Print("\nEnter first calculator number: ")

	calcFirstStr, _ := reader.ReadString('\n')
	calcFirstStr = strings.TrimSpace(calcFirstStr)

	calcFirst, _ := strconv.Atoi(calcFirstStr)

	fmt.Print("Enter operator (+, -, *, /): ")

	operatorStr, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(operatorStr)

	fmt.Print("Enter second calculator number: ")

	calcSecondStr, _ := reader.ReadString('\n')
	calcSecondStr = strings.TrimSpace(calcSecondStr)

	calcSecond, _ := strconv.Atoi(calcSecondStr)

	switch operator {
	case "+":
		fmt.Println("Result:", calcFirst+calcSecond)

	case "-":
		fmt.Println("Result:", calcFirst-calcSecond)

	case "*":
		fmt.Println("Result:", calcFirst*calcSecond)

	case "/":
		if calcSecond == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Println("Result:", calcFirst/calcSecond)
		}

	default:
		fmt.Println("Invalid operator")
	}

	// ========================================================
	// PROGRAM 12: CHECK CHARACTER TYPE
	// ========================================================
	//
	// Demonstrates multiple conditions.
	//
	// This example checks whether the entered character is:
	//
	// vowel
	// consonant
	// or something else
	//
	// ========================================================

	fmt.Print("\nEnter a character: ")

	charStr, _ := reader.ReadString('\n')
	charStr = strings.TrimSpace(charStr)

	if len(charStr) == 1 {
		char := charStr[0]

		if char == 'a' || char == 'e' || char == 'i' ||
			char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' ||
			char == 'O' || char == 'U' {

			fmt.Println("Vowel")

		} else if (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') {

			fmt.Println("Consonant")

		} else {
			fmt.Println("Other character")
		}
	} else {
		fmt.Println("Please enter exactly one character")
	}

	// ========================================================
	// PROGRAM 13: LEAP YEAR
	// ========================================================
	//
	// A year is a leap year when:
	//
	// 1. It is divisible by 400
	//
	// OR
	//
	// 2. It is divisible by 4
	//    AND
	//    it is NOT divisible by 100
	//
	// ========================================================

	fmt.Print("\nEnter a year: ")

	yearStr, _ := reader.ReadString('\n')
	yearStr = strings.TrimSpace(yearStr)

	year, _ := strconv.Atoi(yearStr)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not a leap year")
	}

	// ========================================================
	// PROGRAM 14: LOGIN STATUS
	// ========================================================
	//
	// Demonstrates boolean conditions.
	//
	// ========================================================

	isLoggedIn := true
	isAdmin := false

	if isLoggedIn && isAdmin {
		fmt.Println("\nLogged in as Admin")
	} else if isLoggedIn {
		fmt.Println("\nLogged in as normal user")
	} else {
		fmt.Println("\nUser is not logged in")
	}

	// ========================================================
	// PROGRAM 15: NESTED IF
	// ========================================================
	//
	// An if statement can exist inside another if statement.
	//
	// ========================================================

	passwordCorrect := true
	accountActive := true

	if passwordCorrect {
		if accountActive {
			fmt.Println("Login successful")
		} else {
			fmt.Println("Account is inactive")
		}
	} else {
		fmt.Println("Incorrect password")
	}
}

// ============================================================
// QUICK NOTES
// ============================================================
//
// 1. BASIC IF
//
// if condition {
//     // code
// }
//
//
//
// 2. IF / ELSE
//
// if condition {
//     // true
// } else {
//     // false
// }
//
//
//
// 3. IF / ELSE IF / ELSE
//
// if condition1 {
//     // condition1 is true
// } else if condition2 {
//     // condition2 is true
// } else {
//     // none of the conditions are true
// }
//
//
//
// 4. COMPARISON OPERATORS
//
// ==    Equal to
// !=    Not equal to
// >     Greater than
// <     Less than
// >=    Greater than or equal to
// <=    Less than or equal to
//
//
//
// 5. LOGICAL OPERATORS
//
// &&    AND
// ||    OR
// !     NOT
//
// Example:
//
// if age >= 18 && age <= 60 {
//     fmt.Println("Adult")
// }
//
//
//
// 6. MODULUS OPERATOR
//
// % gives the remainder.
//
// Example:
//
// 10 % 2 -> 0
// 11 % 2 -> 1
//
// This is commonly used to check even/odd numbers.
//
//
//
// 7. SWITCH
//
// switch value {
// case 1:
//     // code
// case 2:
//     // code
// default:
//     // code
// }
//
// Go automatically stops after a matching case.
// You normally do NOT need break.
//
//
//
// 8. VARIABLE DECLARATION
//
// var result string
//
// Creates a variable with its zero value.
//
// For string:
//
// ""
//
// For int:
//
// 0
//
// For bool:
//
// false
//
//
//
// 9. TYPE CONVERSION
//
// User input from ReadString() is a string.
//
// Convert string -> int:
//
// number, err := strconv.Atoi(input)
//
//
//
// 10. TRIM INPUT
//
// strings.TrimSpace(input)
//
// Removes:
//
// - spaces
// - tabs
// - newline characters
//
//
//
// 11. READING TERMINAL INPUT
//
// reader := bufio.NewReader(os.Stdin)
//
// input, err := reader.ReadString('\n')
//
//
//
// IMPORTANT
//
// In real Go programs, don't ignore errors using:
//
// value, _ := ...
//
// The "_" means:
//
// "I don't care about this value."
//
// For learning, it is sometimes used to keep examples simple.
//
// In production code, errors should normally be checked:
//
// value, err := strconv.Atoi(input)
//
// if err != nil {
//     fmt.Println("Invalid input")
// }
//
// ============================================================
//
// CONDITIONS CHEAT SHEET
//
// Check positive:
//
// if number > 0
//
// Check negative:
//
// if number < 0
//
// Check zero:
//
// if number == 0
//
// Check even:
//
// if number%2 == 0
//
// Check odd:
//
// if number%2 != 0
//
// Check range:
//
// if age >= 18 && age <= 60
//
// Check either condition:
//
// if age < 18 || age > 60
//
// Check NOT:
//
// if !isLoggedIn
//
// ============================================================
