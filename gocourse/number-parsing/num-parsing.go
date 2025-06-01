package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
		return
	}
	fmt.Printf("Parsed integer: %v of type: %T", num, num)

	num64, err := strconv.ParseInt(numStr, 10, 64) // Parse as base 10, 64-bit integer
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
		return
	}
	fmt.Println("Parsed int64: ", num64)

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64) // Parse as 64-bit float
	if err != nil {
		fmt.Println("Error parsing the float value: ", err)
		return
	}
	fmt.Printf("Parsed float: %v of type: %T\n", floatVal, floatVal)
}
