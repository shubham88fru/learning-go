package main

func main() {

	increasing := adder()
	println(increasing())
	println(increasing())
	println(increasing())
	println(increasing())

	decreasing := substractor()
	println(decreasing(1))
	println(decreasing(2))
	println(decreasing(3))
	println(decreasing(4))

}

func substractor() func(int) int {
	val := 99
	println("previous value of val is: ", val)
	return func(amt int) int {
		val -= amt
		println("substracted ", amt, " from val")
		return val
	}
}

func adder() func() int {
	i := 0
	println("previous value of i is: ", i)
	return func() int {
		i += 1
		println("added 1 to i")
		return i
	}
}
