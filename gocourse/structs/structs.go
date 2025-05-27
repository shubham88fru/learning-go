package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println("Struct Person: ", p)
	fmt.Println("First Name: ", p.firstName)
	fmt.Println("Last Name: ", p.lastName)
	fmt.Println("Age: ", p.age)

	// anonymous structs
	user := struct {
		userName string
		email    string
	}{
		userName: "johndoe",
		email:    "johndoe@johndoe.com",
	}
	fmt.Println("Anonymous Struct User: ", user)
}
