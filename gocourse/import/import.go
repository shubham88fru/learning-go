package main

import (
	"fmt"
	foo "net/http" // Importing the net/http package with an alias
)

func main() {
	fmt.Println("Hello, go standard library!")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("HTTP Response Status: ", resp.Status)
}
