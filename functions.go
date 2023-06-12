//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting (name string) {
	fmt.Println("Hello", name)
}

func message () string {
	return "This is a test"
}

func addThree (a, b, c int) int {
	return a + b + c
}

func returnNumber () int {
	return 5
}

func returnTwoNumbers () (int, int) {
	return 5, 10
}

func main() {
	greeting("Bob")
	fmt.Println(message())
	a, b := returnTwoNumbers()
	answer := addThree(returnNumber(), a, b)
	fmt.Println(answer)
}


