package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Hello, World!"

	fmt.Println(len(str))
	str1 := "Hello"
	str2 := "World"

	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Printf("%T\n", str[0]) // rune
	fmt.Println(str[1:5])

	// string std library functions

	// 1. string conversion
	num := 18
	str3 := strconv.Itoa(num) // convert int to string
	fmt.Println(str3)

	// 2. string split
	fruits := "apple, orange, banana"
	parts := strings.Split(fruits, ", ")
	fmt.Println(parts)

	// 3. string concatenation
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ",")
	fmt.Println(joined)

	// 4. string contains
	fmt.Println(strings.Contains(str, "World"))
	fmt.Println(strings.Contains(str, "world")) // case-sensitive

	// 5. string replace
	replaced := strings.Replace(str, "World", "Gopher", 1)
	fmt.Println(replaced)

	// 6. trim spaces
	strwspace := "   Hello, World!   "
	trimmed := strings.TrimSpace(strwspace)
	fmt.Println(trimmed)

	// 7. string to lower/upper case
	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Println(lower)
	fmt.Println(upper)

	// 8. count of substring/character/rune
	count := strings.Count(str, "o")
	fmt.Println(count) // 2
	count = strings.Count(str, "l")
	fmt.Println(count) // 3
	count = strings.Count(str, "llo")
	fmt.Println(count) // 1

	// String builder
	var builder strings.Builder // mutable string
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")
	fmt.Println(builder.String())

	builder.WriteRune(' ') // Write a rune
	builder.WriteString("Go is great!")
	fmt.Println(builder.String())

	builder.Reset()               // clear the builder
	fmt.Println(builder.String()) // prints empty string
}
