//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string 

	
func showLine (line []Part) {
	fmt.Println("The line contains:")
	for _, part := range line {
		fmt.Println(part)
	}
	fmt.Println()
}




func main() {
	line := []Part{"wheels", "doors", "hood"}
	showLine(line)
	
	line = append(line, "seats", "dashboard")
	showLine(line)
	
	line = line[3:]
	showLine(line)


}
