package main

func main() {

	var ptr *int
	var a int = 10
	ptr = &a // ptr now points to a
	println("Value of a is: ", a)
	println("Address of a is: ", &a)
	println("Value of ptr is: ", ptr)
	println("Value pointed to by ptr is: ", *ptr)

	modifyValue(ptr) // pass the pointer
	println("Value of a after modification is: ", a)
}

func modifyValue(ptr *int) {
	*ptr += 1
	println("Value pointed to by ptr is: ", *ptr)
}
