// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var str string
	var num int
	var isTrue bool
	// Display the value of those variables.
	fmt.Println(str)
	fmt.Println(num)
	fmt.Println(isTrue)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	greeting := "Hello"
	favNum := 12
	isHappy := true
	// Display the value of those variables.
	fmt.Printf("String variable: %v\n", greeting)
	fmt.Printf("Num variable: %v\n", favNum)
	fmt.Printf("Bool variable: %v\n", isHappy)
	// Perform a type conversion.

	favInt := int(favNum)
	// Display the value of that variable.
	fmt.Printf("Int: %v", favInt)
}
