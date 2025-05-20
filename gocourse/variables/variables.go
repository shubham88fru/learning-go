package main

import "fmt"

var lastName string = "Singh" // Global variable declaration with initialization - package scope

func main() {
	var age int                 // Variable declaration - no initialization
	var name string = "Shubham" // Variable declaration with initialization
	isStudent := true           // Short variable declaration - no type specified, inferred from value

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
}
