package main

import "fmt"

func main() {
	age := 32

	//var agePointer *int
	agePointer := &age  // use `&` in front of a normal variable and get its address.
	fmt.Println("Age: ", *agePointer) // use `*` in front of a pointer to reference and the get value at the mem address.
	/*
	 Note how the `*` operator can be used in two places. First, in front of a type (int, float etc)
	 in which case usually it indicates that the variable on the left is a pointer of type xyz.
	 Second, in front of a pointer varible, in which case, it simply dereferences the value stored
	 at the pointer location.
	*/
	
	
	fmt.Println("Age: ", age)

	adultYears := getAdultYears(age)
	adultYears2 := getAdultYearsPointer(&age)
	fmt.Println(adultYears)
	fmt.Println(adultYears2)

	updateAdultYears(&age)
	fmt.Println("After update:" , age)
}

func getAdultYears(age int) int  {
	return age - 18
}

func getAdultYearsPointer(age *int) int {
	return *age - 18
}

func updateAdultYears(age *int) {
	//directly manipulating the value
	//using pointer, instead of returning.
	*age = *age - 18
}