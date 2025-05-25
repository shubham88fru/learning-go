package main

func main() {
	println(sumOfDigits(54321))
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return (n % 10) + sumOfDigits(n/10)
}
