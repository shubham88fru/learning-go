package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	//Shortcut that can be used to declare and assign a varible, where
	//the type is okay to be infered by go. NOTE: We can leverage this syntax
	//when we explicitly provide type for a variable, in such case we'll have to
	//go with a usual assignment using `var`.
	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/100), years)

	futureRealValue := futureValue/math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}