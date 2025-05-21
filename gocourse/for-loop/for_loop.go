package main

import "fmt"

func main() {

	//simple for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//iterate over a slice
	numbers := []int{1, 2, 3, 4, 5, 6}
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// for loop as while loop - there's no while loop in Go
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
